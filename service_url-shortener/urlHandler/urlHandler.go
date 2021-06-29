// The package Url-handler is providing the REST-endpoint and the Url-redirecting functionalities.
package urlHandler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Imports:

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	shortener "urlShortener/urlShortener"
	store "urlShortener/urlStorage"

	"github.com/gin-gonic/gin"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

type routerEngine struct {
	router *gin.Engine
}

var (
	routerService   = &routerEngine{}
)
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Constants:

// The GET "/" (default route) lands the Demo frontpage.
const GET_Entry = "/"

// The POST "/createShortUrl" for creating a short-link.
const POST_CreateURL = "/createShortUrl"

// The GET "/:shortUrl" for redirecting the short-link to the original URL.
const GET_ShortUrl = "/:shortUrl"

// The GET "/shortUrlStats/:shortUrl" for requesting the short-link meta data.
const GET_ShortUrlStats = "/shortUrlStats/:shortUrl"

// The POST "/shortUrlDelete/:shortUrl" for requesting the short-link to be deleted.
const POST_ShortUrlDelete = "/shortUrlDelete"

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exported:

// Definition of the request model for short-link creation
type UrlCreationRequest struct {
    LongUrl string `json:"long_url" binding:"required"`
}

// Definition of the response model for Create-Request
type Response struct {
    Status int `json:"status" binding:"required"`
    Message string `json:"message" binding:"required"`
    Response interface{} `json:"response" binding:"required"`
}

// Definition of the response model data for Create-Request
type ResponseUrls struct {
    LongUrl string `json:"long_url" binding:"required"`
    ShortUrl string `json:"short_url" binding:"required"`
}

// Definition of the response model data for Statistics Request
type ResponseStats struct {
    UrlCode string `json:"url_code" binding:"required"`
    VisitCount int `json:"visit_count" binding:"required"`
}

// Definition of the request model for short-link deletion
type UrlDeleteRequest struct {
    ShortUrl string `json:"short_url" binding:"required"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Implementation:

// Initialize the REST-
func Initialize(routeIp string, routePort string) (err error) {

	routerService.router = gin.Default()

	routerService.router.GET(GET_Entry, func(ctx *gin.Context) {
		GetEntry(ctx)
	})

	routerService.router.POST(POST_CreateURL, func(ctx *gin.Context) {
		ShortUrlCreate(ctx, routeIp+":"+routePort)
	})

	routerService.router.GET(GET_ShortUrl, func(ctx *gin.Context) {
		ShortUrlRedirect(ctx)
	})

	routerService.router.GET(GET_ShortUrlStats, func(ctx *gin.Context) {
		ShortUrlStatistics(ctx)
	})

	routerService.router.POST(POST_ShortUrlDelete, func(ctx *gin.Context) {
		ShortUrlDelete(ctx, routeIp+":"+routePort)
	})

	err = routerService.router.Run(":" + routePort)
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

    return err
}


func GetEntry(ctx *gin.Context) {
	demoContent, err := ioutil.ReadFile("static-data/index.html")
    if err == nil {
        ctx.Data(http.StatusOK, "text/html; charset=utf-8", demoContent)
    } else {
        errorResponse(ctx, http.StatusInternalServerError, err)
    }
}

func ShortUrlCreate(ctx *gin.Context, urlCodeHostname string) {
    var creationRequest UrlCreationRequest
    if err :=ctx.ShouldBindJSON(&creationRequest); err != nil {
       ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    longUrl := parseUrl(creationRequest.LongUrl)

    var err error
    urlRecord := store.NewUrlItem()
    urlRecord.UrlBase = longUrl
    urlRecord.UrlCode, urlRecord.Id, err = shortener.GetUrlEncoding(longUrl)

    if err == nil {
        if ! store.UrlRecordIsExisting(*urlRecord) {
            store.UrlRecordWrite(*urlRecord)
        }

        // send response
       ctx.Header("Content-Type", "application/json")
        responseUrls := &Response{
            Status:  http.StatusOK,
            Message: "Short url created successfully.",
            Response: ResponseUrls{
                LongUrl:  longUrl,
                ShortUrl: urlCodeHostname + "/" + urlRecord.UrlCode,
            },
        }
       ctx.JSON(http.StatusOK, responseUrls)
    } else {
        errorResponse(ctx, http.StatusInternalServerError, err)
    }
}

func ShortUrlRedirect(ctx *gin.Context) {
    shortUrl :=ctx.Param("shortUrl")
    urlRecord, err := store.UrlRecordRead(shortUrl)

    if err == nil {
        // count the redirect and write back to store
        urlRecord.VisitCount += 1
        store.UrlRecordWrite(urlRecord)

        ctx.Redirect(http.StatusFound, urlRecord.UrlBase)
    } else {
        errorResponse(ctx, http.StatusInternalServerError, err)
    }
}

func ShortUrlStatistics(ctx *gin.Context) {
    shortUrl :=ctx.Param("shortUrl")
    urlRecord, err := store.UrlRecordRead(shortUrl)
    statusCode := http.StatusFound

    if err == nil {
       ctx.Header("Content-Type", "application/json")
        responseUrls := &Response{
            Status:  statusCode,
            Message: "Short url found.",
            Response: ResponseStats{
                UrlCode: urlRecord.UrlCode,
                VisitCount: int(urlRecord.VisitCount),
            },
        }
       ctx.JSON(statusCode, responseUrls)
    } else {
        errorResponse(ctx, http.StatusInternalServerError, err)
    }
}

func ShortUrlDelete(ctx *gin.Context, urlCodeHostname string) {
    var deleteRequest UrlDeleteRequest
    if err :=ctx.ShouldBindJSON(&deleteRequest); err != nil {
       ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    shortUrl := strings.Replace(deleteRequest.ShortUrl, urlCodeHostname + "/", "", -1)
    urlRecord, err := store.UrlRecordRead(shortUrl)
    statusCode := http.StatusFound

    var errs error = nil
    if err == nil {
        if store.UrlRecordIsExisting(urlRecord) {
           errs = store.UrlRecordDelete(urlRecord.UrlCode)
        }
        if errs == nil {
            // send response
            ctx.Header("Content-Type", "application/json")
            responseUrls := &Response{
                Status:  statusCode,
                Message: "Short url deleted successfully.",
                Response: gin.H{
                    "data": "",
                },        
            }
            ctx.JSON(statusCode, responseUrls)
        }
    }

    if err != nil || errs != nil {
        errorResponse(ctx, http.StatusInternalServerError, err)
    }
}
////////////////////////////////////////////////////////////

// Function to parse URLs which might be recognized as Request-URIs
// Shall return an absolute URI as valid URL
func parseUrl(inputUrl string) string {
    var err error
    // force interpretation as absolute URI, see https://golang.org/pkg/net/url/#ParseRequestURI
    u, err := url.ParseRequestURI(inputUrl)
    if err != nil || u.Host == "" {
        u, repErr := url.ParseRequestURI("https://" + inputUrl)
        if repErr != nil {
            fmt.Printf("Could not parse raw url: %s, error: %v", inputUrl, err)
            return u.String()
        }
        err = nil
        return u.String()
    }
    return u.String()
}

// Function for a general error response to requests
func errorResponse(ctx *gin.Context, errorCode int, err error) {
   ctx.Header("Content-Type", "application/json")
   ctx.JSON(errorCode, gin.H{
        "Status":  errorCode,
        "Message": err,
        "Response": gin.H{
            "data": "",
        },
    })
}

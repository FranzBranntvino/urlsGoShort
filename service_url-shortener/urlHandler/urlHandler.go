// The package Url-handler is providing the REST-endpoint and the Url-redirecting functionalities.
package urlHandler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Imports:

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	shortener "urlShortener/urlShortener"
	store "urlShortener/urlStorage"

	"github.com/gin-gonic/gin"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

type routerEngine struct {
	router *gin.Engine
    ip string
    port string
}

var (
	routerService = &routerEngine{nil, "", ""}
)
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Constants:

// The GET "/" (default route) lands the Demo frontpage.
const GET_entry = "/"

// The POST "/createShortUrl" for creating a short-link.
const POST_createURL = "/createShortUrl"

// The GET "/:shortUrl" for redirecting the short-link to the original URL.
const GET_shortUrl = "/:shortUrl"

// The GET "/shortUrlStats/:shortUrl" for requesting the short-link meta data.
const GET_shortUrlStats = "/shortUrlStats/:shortUrl"

// The DELETE "/shortUrlDelete/:shortUrl" for requesting the short-link to be deleted.
const DELETE_shortUrlDelete = "/shortUrlDelete/:shortUrl"

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

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Implementation:

// Initialize the REST-Service
func Initialize(routeIp string, routePort string) {
	routerService.router = gin.Default()
    routerService.ip = routeIp
    routerService.port = routePort

    routerService.router.Static("/static-data", "./static-data")
    // routerService.router.StaticFS("/static-data", http.Dir("static_data"))
    routerService.router.StaticFile("/index.html", "./static-data/index.html")
    routerService.router.StaticFile("/favicon.ico", "./static-data/favicon.ico")

	routerService.router.GET(GET_entry, func(ctx *gin.Context) {
		getEntry(ctx)
	})

	routerService.router.POST(POST_createURL, func(ctx *gin.Context) {
		shortUrlCreate(ctx)
	})

	routerService.router.GET(GET_shortUrl, func(ctx *gin.Context) {
		shortUrlRedirect(ctx)
	})

	routerService.router.GET(GET_shortUrlStats, func(ctx *gin.Context) {
		shortUrlStatistics(ctx)
	})

    routerService.router.DELETE(DELETE_shortUrlDelete, func(ctx *gin.Context) {
		shortUrlDelete(ctx)
	})
}

// Starting the router service
// To be called only after initializing function
func StartServing() (error) {
    if routerService.router != nil {
        err := routerService.router.Run(":" + routerService.port)
        if err != nil {
            panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
        }
        return nil
    }
    return errors.New("router not initialized")
}

// Stopping the router service
func StopServing() (error) {
    if routerService.router != nil {
        // TODO implement shutdown here ...
        return nil
    }
    return errors.New("router not initialized")
}

////////////////////////////////////////////////////////////

// Handling static Index
func getEntry(ctx *gin.Context) {
	demoContent, err := ioutil.ReadFile("static-data/index.html")
    if err == nil {
        ctx.Data(http.StatusOK, "text/html; charset=utf-8", demoContent)
    } else {
        errorResponse(ctx, http.StatusInternalServerError, err)
    }
}

// Handling the Request about creation for a short link and it's DB record entry
func shortUrlCreate(ctx *gin.Context) {
    var urlCodeHostname string = routerService.ip+":"+routerService.port
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
            Status:  http.StatusCreated,
            Message: "Short url created successfully.",
            Response: ResponseUrls{
                LongUrl:  longUrl,
                ShortUrl: urlCodeHostname + "/" + urlRecord.UrlCode,
            },
        }
       ctx.JSON(http.StatusCreated, responseUrls)
    } else {
        errorResponse(ctx, http.StatusInternalServerError, err)
    }
}

// Handling the Request about redirection/routing for a short link to it's original location
func shortUrlRedirect(ctx *gin.Context) {
    shortUrl :=ctx.Param("shortUrl")
    urlRecord, err := store.UrlRecordRead(shortUrl)
    if err == nil {
        // count the redirect and write back to store
        urlRecord.VisitCount += 1
        store.UrlRecordWrite(urlRecord)
        ctx.Redirect(http.StatusFound, urlRecord.UrlBase)
    } else {
        errorResponse(ctx, http.StatusBadRequest, err)
    }
}

// Handling the Request about the statistics for a short link redirecting to it's original location
func shortUrlStatistics(ctx *gin.Context) {
    shortUrl :=ctx.Param("shortUrl")
    urlRecord, err := store.UrlRecordRead(shortUrl)
    statusCode := http.StatusOK

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

// Handling the Request about deletion for a short link and it's DB record entry
func shortUrlDelete(ctx *gin.Context) {
    shortUrl := ctx.Params.ByName("shortUrl")
    urlRecord, err := store.UrlRecordRead(shortUrl)
    statusCode := http.StatusNoContent

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
        errorResponse(ctx, http.StatusNotFound, err)
    }
}

// Function to parse URLs which might be recognized as Request-URIs
// Shall return an absolute URI as valid URL
func parseUrl(inputUrl string) string {
    var err error
    u, err := url.Parse(inputUrl)
    // force interpretation as absolute URI, see https://golang.org/pkg/net/url/#ParseRequestURI
    if (err != nil || u.Host == "") && (u.Scheme != "http" && u.Scheme != "https" && u.Scheme != "file") {
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

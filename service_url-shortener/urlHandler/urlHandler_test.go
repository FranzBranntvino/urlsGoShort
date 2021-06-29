package urlHandler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Imports:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Constants:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exported:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Implementation:

func TestGetEntry(t *testing.T) {
	_, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	status := rr.Code
	assert.Equal(t, http.StatusOK, status)
}

// curl --request POST --data '{ "long_url": "https://github.com/FranzBranntvino/urlsGoShort/" }' http://localhost:8080/createShortUrl && echo " "
func TestShortUrlCreate(t *testing.T) {
	var hostStr string = "http://172.17.0.1:8080"
	var jsonStr = []byte(`{"long_url": "https://github.com/FranzBranntvino/urlsGoShort/"}`)
	responseBody := bytes.NewBuffer(jsonStr)
	resp, err := http.Post(hostStr+POST_CreateURL, "application/json", responseBody)

	if err == nil {
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		expected := `{"status":200,"message":"Short url created successfully.","response":{"long_url":"https://github.com/FranzBranntvino/urlsGoShort/","short_url":"localhost:8080/JF95kx753gh"}}`

		fmt.Println("Returned: " + string(bodyBytes))

		assert.Equal(t, true, strings.Contains(string(bodyBytes), expected))
	}
	assert.Equal(t, nil, err)
}

// curl http://localhost:8080/JF95kx753gh && echo " "
func TestShortUrlRedirect(t *testing.T) {
	resp, err := http.Get("http://172.17.0.1:8080/JF95kx753gh")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	stringBody := string(body)
	var rex = regexp.MustCompile(`body|html`)
	assert.Equal(t, true, rex.MatchString(stringBody))
}

func TestShortUrlDelete(t *testing.T) {
	// curl --request POST --data '{ "long_url": "https://github.com/FranzBranntvino/urlsGoShort/" }' http://localhost:8080/createShortUrl && echo " "
	// {"status":200,"message":"Short url created successfully.","response":{"long_url":"https://github.com/FranzBranntvino/urlsGoShort/","short_url":"localhost:8080/JF95kx753gh"}}
	// curl --request POST --data '{ "short_url": "localhost:8080/JF95kx753gh" }' http://localhost:8080/shortUrlDelete && echo " "
	// {"status":302,"message":"Short url deleted successfully.","response":{"data":""}}
}

func TestShortUrlStats(t *testing.T) {
	// curl --request POST --data '{ "long_url": "https://github.com/FranzBranntvino/urlsGoShort/" }' http://localhost:8080/createShortUrl && echo " "
	// {"status":200,"message":"Short url created successfully.","response":{"long_url":"https://github.com/FranzBranntvino/urlsGoShort/","short_url":"localhost:8080/JF95kx753gh"}}
	// curl http://localhost:8080/JF95kx753gh
	// <a href="https://github.com/FranzBranntvino/urlsGoShort/">Found</a>.
	// curl http://localhost:8080/JF95kx753gh
	// <a href="https://github.com/FranzBranntvino/urlsGoShort/">Found</a>.
	// curl http://localhost:8080/JF95kx753gh
	// <a href="https://github.com/FranzBranntvino/urlsGoShort/">Found</a>.
	// curl http://localhost:8080/shortUrlStats/JF95kx753gh && echo " "
	// {"status":302,"message":"Short url found.","response":{"url_code":"JF95kx753gh","visit_count":3}}
}

////////////////////////////////////////////////////////////

func TestParseUrl(t *testing.T) {
	assert := assert.New(t)

	rString := parseUrl("https://github.com/FranzBranntvino/urlsGoShort")
	eString := "https://github.com/FranzBranntvino/urlsGoShort"
	assert.Equal(eString, rString)

	rString = parseUrl("www.google.de")
	eString = "https://www.google.de"
	assert.Equal(eString, rString)

	rString = parseUrl("google.de")
	eString = "https://google.de"
	assert.Equal(eString, rString)
}

package urlHandler

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Imports:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Global:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Constants:

const (
	hostStr      = "http://172.17.0.1:" + routePort
	routeAddress = "localhost"
	routePort    = "8080"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Exported:

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Implementation:

func init() {
	Initialize(routeAddress, routePort)
}

func Test_Initialize(t *testing.T) {
	respRec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/WhateverIsNotWorking", nil)
	routerService.router.ServeHTTP(respRec, req)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusInternalServerError, respRec.Code)
}

func Test_StartServing(t *testing.T) {
	// assert.Equal(t, nil, StartServing())
}

func Test_StopServing(t *testing.T) {
	assert.Equal(t, nil, StopServing())
}

func Test_getEntry(t *testing.T) {
	// // mock the server ...
	// mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	// mock a response ...
	// 	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	// 	fmt.Fprintln(w, "This is a super simple, by default html, response")
	// }))
	// defer mockServer.Close()
	// // request front page ...
	// _, err := http.NewRequest(http.MethodGet, mockServer.URL+"/", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// respRec := httptest.NewRecorder()
	// req, err := http.NewRequest(http.MethodGet, "/", nil)
	// routerService.router.ServeHTTP(respRec, req)
	// status := respRec.Code
	// assert.Equal(t, http.StatusOK, status)

	resp, err := http.Get(hostStr + GET_entry)
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	assert.Equal(t, nil, err)
}

func Test_shortUrlCreate(t *testing.T) {
	// curl --request POST --data '{ "long_url": "https://github.com/FranzBranntvino/urlsGoShort/" }' http://localhost:8080/createShortUrl && echo " "
	var jsonBytes = []byte(`{"long_url": "https://github.com/FranzBranntvino/urlsGoShort/"}`)
	responseBody := bytes.NewBuffer(jsonBytes)
	resp, err := http.Post(hostStr+POST_createURL, "application/json", responseBody)
	// {"status":201,"message":"Short url created successfully.","response":{"long_url":"https://github.com/FranzBranntvino/urlsGoShort/","short_url":"localhost:8080/JF95kx753gh"}}
	expected := `{"status":201,"message":"Short url created successfully.","response":{"long_url":"https://github.com/FranzBranntvino/urlsGoShort/","short_url":"localhost:8080/JF95kx753gh"}}`

	if err == nil {
		defer resp.Body.Close()
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Returned: " + string(bodyBytes))
		assert.Equal(t, true, strings.Contains(string(bodyBytes), expected))
	}
	assert.Equal(t, nil, err)
}

func Test_shortUrlRedirect(t *testing.T) {
	// Fail
	resp, err := http.Get(hostStr + "/0815abcdefg")
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Success
	// curl --request POST --data '{ "long_url": "https://github.com/FranzBranntvino/urlsGoShort/" }' http://localhost:8080/createShortUrl && echo " "
	var jsonStr = []byte(`{"long_url": "https://github.com/FranzBranntvino/urlsGoShort/"}`)
	responseBody := bytes.NewBuffer(jsonStr)
	resp, err = http.Post(hostStr+POST_createURL, "application/json", responseBody)
	if err == nil {
		defer resp.Body.Close()
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	}
	assert.Equal(t, nil, err)

	// curl http://localhost:8080/JF95kx753gh && echo " "
	resp, err = http.Get(hostStr + "/JF95kx753gh")
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

func Test_shortUrlDelete(t *testing.T) {
	// curl --request POST --data '{ "long_url": "https://github.com/FranzBranntvino/urlsGoShort/" }' http://localhost:8080/createShortUrl && echo " "
	// {"status":201,"message":"Short url created successfully.","response":{"long_url":"https://github.com/FranzBranntvino/urlsGoShort/","short_url":"localhost:8080/JF95kx753gh"}}
	var jsonStr = []byte(`{"long_url": "https://github.com/FranzBranntvino/urlsGoShort/"}`)
	responseBody := bytes.NewBuffer(jsonStr)
	resp, err := http.Post(hostStr+POST_createURL, "application/json", responseBody)

	if err == nil {
		defer resp.Body.Close()
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		expected := `{"status":201,"message":"Short url created successfully.","response":{"long_url":"https://github.com/FranzBranntvino/urlsGoShort/","short_url":"localhost:8080/JF95kx753gh"}}`

		fmt.Println("Returned: " + string(bodyBytes))

		assert.Equal(t, true, strings.Contains(string(bodyBytes), expected))
	}
	assert.Equal(t, nil, err)

	// curl -i -X DELETE http://localhost:8080/shortUrlDelete/JF95kx753gh && echo " "
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", hostStr+"/shortUrlDelete/JF95kx753gh", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(t, nil, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode)

	fmt.Println("Returned: ", string(respBody))
	fmt.Println("Status: ", resp.Status)
	fmt.Println("Header: ", resp.Header)
}

func Test_shortUrlStats(t *testing.T) {
	// curl --request POST --data '{ "long_url": "https://github.com/FranzBranntvino/urlsGoShort/" }' http://localhost:8080/createShortUrl && echo " "
	// {"status":201,"message":"Short url created successfully.","response":{"long_url":"https://github.com/FranzBranntvino/urlsGoShort/","short_url":"localhost:8080/JF95kx753gh"}}
	// curl http://localhost:8080/JF95kx753gh
	// <a href="https://github.com/FranzBranntvino/urlsGoShort/">Found</a>.
	// curl http://localhost:8080/JF95kx753gh
	// <a href="https://github.com/FranzBranntvino/urlsGoShort/">Found</a>.
	// curl http://localhost:8080/JF95kx753gh
	// <a href="https://github.com/FranzBranntvino/urlsGoShort/">Found</a>.
	// curl http://localhost:8080/shortUrlStats/JF95kx753gh && echo " "
	// {"status":200,"message":"Short url found.","response":{"url_code":"JF95kx753gh","visit_count":3}}

	// curl --request POST --data '{ "long_url": "https://github.com/FranzBranntvino/urlsGoShort/" }' http://localhost:8080/createShortUrl && echo " "
	var jsonStr = []byte(`{"long_url": "https://github.com/FranzBranntvino/urlsGoShort/"}`)
	responseBody := bytes.NewBuffer(jsonStr)
	resp, err := http.Post(hostStr+POST_createURL, "application/json", responseBody)
	if err == nil {
		defer resp.Body.Close()
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	}
	assert.Equal(t, nil, err)

	// curl http://localhost:8080/JF95kx753gh && echo " "
	_, err = http.Get(hostStr + "/JF95kx753gh")
	assert.Equal(t, nil, err)
	_, err = http.Get(hostStr + "/JF95kx753gh")
	assert.Equal(t, nil, err)
	_, err = http.Get(hostStr + "/JF95kx753gh")
	assert.Equal(t, nil, err)

	// request stats data ...
	resp, err = http.Get(hostStr + "/shortUrlStats/JF95kx753gh")
	if err == nil {
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		expected := `{"status":200,"message":"Short url found.","response":{"url_code":"JF95kx753gh","visit_count":3}}`
		// expected := `{"status":200,"message":"Short url found.","response":{"url_code":"JF95kx753gh","visit_count":`
		assert.Equal(t, true, strings.Contains(string(bodyBytes), expected))

		fmt.Println("Returned: " + string(bodyBytes))
	}
	assert.Equal(t, nil, err)
}

////////////////////////////////////////////////////////////

func Test_parseUrl(t *testing.T) {
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

	rString = parseUrl("https://godbolt.org/#z:OYLghAFBqd5QCxAYwPYBMCmBRdBLAF1QCcAaPECAM1QDsCBlZAQwBtMQBGAFlJvoCqAZ0wAFAB4gA5AAYppAFZdSrZrVDIApACYAQjt2kR7ZATx1KmWugDCqVgFcAtrWVX0AGTy1MAOWcARpjEIADMAJykAA6oQoTmtHaOLsoxcWZ0Xj7%2BTkEhEUaYJhm0DATMxARJzq6chcUJZRUEWX6BwWGRQuWV1Sl13c2tOXmdAJRGqA7EyBxSOqHeyI5YANSaoTYVxMwAnhvYmjIAggtLK5jrm%2BbdxJjMTgdHp9qLtMsOaxs2yLfewE8Ts9ngB6ABUqzKxAcplWWCo3nidCEqzBIOeBF2UUw8NWtxhBFWAH0hKgnJgicwHEQicArMFmAQcUS0F8AOz6E6rbmrAKoeyrKLEVDYyq7GQbTnHHm8/msQXC0WYziS54yvkCoUi4KY7Sqrk8jXyrVK3ahfXSw1yhXasXcC3q60mnW7ACsFs0bIAIniyRSqTS6T4dkz0CyMJgPScQTHVgBxekhnGrNh4ZgomjEVYAFSxkZODji6lWokVOrwmCEACUivcRFcfaTyZTqahaYnGczWZHQlLnoMzMgS2XKhXq7X05cAG5sByYXQOKhUYIN9YctU8qhsESkDfcresHd71YHo8G/fbzC788ny/Xy0Xw9X4HeqPHJlOKKqJnfZbplHZoCxxTqgeDoKsqgAF67HYn4VJgUIEhAaC0N0OY6AAbKsM6sGMa5SrGPIDngQ7Zths7zouy7EA6PJ4FQqwQIx3ToCAIDkk4yCfhAdzeEyxBCpgBAsumBC/ggFSrCh3RggcECYThc4LkuwQTKsvH0MEgnCSw3TiZJ0kELJoTYPJ2gYThalxJBmCoFQEDZmMeFgGAGw%2BjIeF4Z6UoyjKLFsWg1JXDY3zrNo2irL4ADyZGYAAjg4bBhRF3yhf5IDuKwtG%2BeRjiUSpWZubl2Vrj6RT1t5x45dy6WBYSqWbMlqzxYl8o6ClmxpQQrEZdYWW9senpei%2Bw1AmNxzEUOhmYOIQq%2Bs2AZtkGDKhuGawANalraZiVjW7CTpwq6VTep7Pidd7Hqd94yldl0XedT73kNb79uUg5SciTKzVmTb%2Bq27bBp2Ybdqsm0jjt477SIHU%2BsdD63o9d2Iw9Z7wwQ0JnfDt0o5jz0DeNzx8asTjMN4EBeeuN61VM9WdY11m2fZin5dRXl0zYeLdWxmVvjKUEwWSUTwYhpgQGD21jntdaYJwYwlfzsFC3cIsEGLW1KpLE7Q3L%2BOnDedwENMtCrBKuvPeNUgTKw0iuvIriyPIqDSCFegGL60yzGFoScPIBDSHITmkOtIDcDIAB0oQyAAHNobJR2y3A8JwrrhOZKjSNw8hOFwMgyKQ9tyKQTtSPIQggHnfsOxMcCwDAiAgDTUTUuQlBoJ%2BeDsCEdQ4oQJDKPwjAsOwXC8APwhiJIDukAA7jsUTSD7Vs23b/uO9IUXUk3hJ2asxxVgAsqswDIEO4Rh9oYeHRAtB0JgeEQIrHcrgssu%2B6vgfB9oUdhzIoRsqEUdvZsjZK6NkudU7pykJnUg2dOBn1zpwaOoC2QYVCJHGQKD86ryLtIUu5dSCVwDqQGu9cZqYGQK2MgFAIAVGAEIUQagigMAQKgae9sfakDblEDujIEgMJ8KwZhrCC7yC4U/EIwBEGhE4YLcRG9GAsLYdgshyBjjEDobg0gKiyj4HtvIAeTA2AcB4HwOgBBx4SGwUoOoqh1AoAMAYFQeAAhl0gBMEUJQy4l0mB7YxRhureH4UwxR7D5Cz2YPPKQi9IG2ywVPYu4go4YQALQYW4EfE%2BqxEER1WNgcQ5DKGMVwL3QqrxZarBdvoPQb8q4TAQPcLAIRyaQOgdnKO4dwjhCjuEHg4VQgYQwp06RIicHeLLhXd%2BExP7cDDq6GQczwqJyTvHTgUdIGhBXvEzRhDLbELrhAJAYjO4t2QrIzuIAKhOFTnUBErB%2BJlwgAEbBARvAVF2AvURZJyT0CirQVg7yp5YBJuoYegK8B3FMHgKclZlH5IoUyD55BNLWynqwZxOxiACywNg9GeBs5RN2QYoexjR5mIsZPQu1iVBqA0A4vQTiXHwHcVETx0hklRQ2fUchJRLDWD6LULR1hhjtC7tEWISJEj2BqKkcVJRhW5A6AMWskK6BNF6FK/oXKVWlB6C0bwbQFWisGOq5IArjV6uyCKrgExSS%2BOtTEzZhdi570PsfU%2B59L6MRvj4e%2BxSiClO9qQCpZzn5lLwpUgwNSA51IaR0ZpwdQiukvkkwZUco7cDZJwDCrpQjcAwi0x1a8xn4J2YHFF2hC2jKjZbCY0LiBxAsNwIAA%3D%3D")
	eString = "https://godbolt.org/#z:OYLghAFBqd5QCxAYwPYBMCmBRdBLAF1QCcAaPECAM1QDsCBlZAQwBtMQBGAFlJvoCqAZ0wAFAB4gA5AAYppAFZdSrZrVDIApACYAQjt2kR7ZATx1KmWugDCqVgFcAtrWVX0AGTy1MAOWcARpjEIADMAJykAA6oQoTmtHaOLsoxcWZ0Xj7%2BTkEhEUaYJhm0DATMxARJzq6chcUJZRUEWX6BwWGRQuWV1Sl13c2tOXmdAJRGqA7EyBxSOqHeyI5YANSaoTYVxMwAnhvYmjIAggtLK5jrm%2BbdxJjMTgdHp9qLtMsOaxs2yLfewE8Ts9ngB6ABUqzKxAcplWWCo3nidCEqzBIOeBF2UUw8NWtxhBFWAH0hKgnJgicwHEQicArMFmAQcUS0F8AOz6E6rbmrAKoeyrKLEVDYyq7GQbTnHHm8/msQXC0WYziS54yvkCoUi4KY7Sqrk8jXyrVK3ahfXSw1yhXasXcC3q60mnW7ACsFs0bIAIniyRSqTS6T4dkz0CyMJgPScQTHVgBxekhnGrNh4ZgomjEVYAFSxkZODji6lWokVOrwmCEACUivcRFcfaTyZTqahaYnGczWZHQlLnoMzMgS2XKhXq7X05cAG5sByYXQOKhUYIN9YctU8qhsESkDfcresHd71YHo8G/fbzC788ny/Xy0Xw9X4HeqPHJlOKKqJnfZbplHZoCxxTqgeDoKsqgAF67HYn4VJgUIEhAaC0N0OY6AAbKsM6sGMa5SrGPIDngQ7Zths7zouy7EA6PJ4FQqwQIx3ToCAIDkk4yCfhAdzeEyxBCpgBAsumBC/ggFSrCh3RggcECYThc4LkuwQTKsvH0MEgnCSw3TiZJ0kELJoTYPJ2gYThalxJBmCoFQEDZmMeFgGAGw%2BjIeF4Z6UoyjKLFsWg1JXDY3zrNo2irL4ADyZGYAAjg4bBhRF3yhf5IDuKwtG%2BeRjiUSpWZubl2Vrj6RT1t5x45dy6WBYSqWbMlqzxYl8o6ClmxpQQrEZdYWW9senpei%2Bw1AmNxzEUOhmYOIQq%2Bs2AZtkGDKhuGawANalraZiVjW7CTpwq6VTep7Pidd7Hqd94yldl0XedT73kNb79uUg5SciTKzVmTb%2Bq27bBp2Ybdqsm0jjt477SIHU%2BsdD63o9d2Iw9Z7wwQ0JnfDt0o5jz0DeNzx8asTjMN4EBeeuN61VM9WdY11m2fZin5dRXl0zYeLdWxmVvjKUEwWSUTwYhpgQGD21jntdaYJwYwlfzsFC3cIsEGLW1KpLE7Q3L%2BOnDedwENMtCrBKuvPeNUgTKw0iuvIriyPIqDSCFegGL60yzGFoScPIBDSHITmkOtIDcDIAB0oQyAAHNobJR2y3A8JwrrhOZKjSNw8hOFwMgyKQ9tyKQTtSPIQggHnfsOxMcCwDAiAgDTUTUuQlBoJ%2BeDsCEdQ4oQJDKPwjAsOwXC8APwhiJIDukAA7jsUTSD7Vs23b/uO9IUXUk3hJ2asxxVgAsqswDIEO4Rh9oYeHRAtB0JgeEQIrHcrgssu%2B6vgfB9oUdhzIoRsqEUdvZsjZK6NkudU7pykJnUg2dOBn1zpwaOoC2QYVCJHGQKD86ryLtIUu5dSCVwDqQGu9cZqYGQK2MgFAIAVGAEIUQagigMAQKgae9sfakDblEDujIEgMJ8KwZhrCC7yC4U/EIwBEGhE4YLcRG9GAsLYdgshyBjjEDobg0gKiyj4HtvIAeTA2AcB4HwOgBBx4SGwUoOoqh1AoAMAYFQeAAhl0gBMEUJQy4l0mB7YxRhureH4UwxR7D5Cz2YPPKQi9IG2ywVPYu4go4YQALQYW4EfE%2BqxEER1WNgcQ5DKGMVwL3QqrxZarBdvoPQb8q4TAQPcLAIRyaQOgdnKO4dwjhCjuEHg4VQgYQwp06RIicHeLLhXd%2BExP7cDDq6GQczwqJyTvHTgUdIGhBXvEzRhDLbELrhAJAYjO4t2QrIzuIAKhOFTnUBErB%2BJlwgAEbBARvAVF2AvURZJyT0CirQVg7yp5YBJuoYegK8B3FMHgKclZlH5IoUyD55BNLWynqwZxOxiACywNg9GeBs5RN2QYoexjR5mIsZPQu1iVBqA0A4vQTiXHwHcVETx0hklRQ2fUchJRLDWD6LULR1hhjtC7tEWISJEj2BqKkcVJRhW5A6AMWskK6BNF6FK/oXKVWlB6C0bwbQFWisGOq5IArjV6uyCKrgExSS%2BOtTEzZhdi570PsfU%2B59L6MRvj4e%2BxSiClO9qQCpZzn5lLwpUgwNSA51IaR0ZpwdQiukvkkwZUco7cDZJwDCrpQjcAwi0x1a8xn4J2YHFF2hC2jKjZbCY0LiBxAsNwIAA%3D%3D"
	assert.Equal(eString, rString)

	rString = parseUrl("")
	eString = "https:"
	assert.Equal(eString, rString)

	rString = parseUrl("file:///home/vagrant/urlsGoShort/.devcontainer/README.md")
	eString = "file:///home/vagrant/urlsGoShort/.devcontainer/README.md"
	assert.Equal(eString, rString)
}

func Test_errorResponse(t *testing.T) {
	routerService.router.GET("/error", func(ctx *gin.Context) {
		err := errors.New("Test-Error.")
		errorResponse(ctx, http.StatusTeapot, err)
		shortUrlStatistics(ctx)
	})
	// TODO create senseful Test for errorResponse
}

package utility

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

const CATAAS_BASEURL = "https://cataas.com/api"

func TestBasicGetResponse(t *testing.T) {
	r := basicGetResponse()
	fmt.Println(r)
}

func TestBuildQueryParam(t *testing.T) {
	qp := make(map[string]string)
	qp["type"] = "cute"
	qp["kind"] = "international"
	absUrl := buildQueryParam(CATAAS_BASEURL+"/cats", qp)
	fmt.Println(absUrl)
}

func basicGetResponse() string {
	// create the http client object
	client := &http.Client{}
	// initialize the query params variable
	queryParams := make(map[string]string)
	// add the values in the query params variable
	queryParams["tags"] = "cute"
	absUrl := buildQueryParam(CATAAS_BASEURL+"/cats", queryParams)
	if req, err := http.NewRequest("GET", absUrl, nil); err != nil {
		panic("Error building request: " + err.Error())
	} else {
		if response, err := client.Do(req); err != nil {
			panic("Error getting response: " + err.Error())
		} else {
			defer response.Body.Close()
			if bodyBytes, err := io.ReadAll(response.Body); err != nil {
				panic("Error reading response body: " + err.Error())
			} else {
				return string(bodyBytes)
			}
		}
	}
}

func buildQueryParam(url string, queryParams map[string]string) string {
	count := 0
	for key, value := range queryParams {
		if count > 0 {
			url = url + "&" + key + "=" + value
		} else {
			url = url + "?" + key + "=" + value
		}
		count++
	}
	return url
}

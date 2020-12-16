package httphelper

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// A simple helper function to pass the correct headers and parse the response for the three different API endpoints that we have to hit.
func GetBytes(token, endpoint string) []byte {
	reqUrl, _ := url.Parse(endpoint)

	tokenString := fmt.Sprintf("Bearer %s", token)

	req := &http.Request{
		URL: reqUrl,
		Header: map[string][]string{
			"Authorization": {tokenString},
			"Content-type":  {"application/json"},
		},
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return bytes
}

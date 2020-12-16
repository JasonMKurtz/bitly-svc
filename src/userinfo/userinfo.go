package userinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var endpoint = "https://api-ssl.bitly.com/v4/user"

type UserinfoEmail struct {
	Email    string `json:"email"`
	Primary  bool   `json:"is_primary"`
	Verified bool   `json:"is_verified"`
}

type Userinfo struct {
	Created   string          `json:"created"`
	Modified  string          `json:"modified"`
	Login     string          `json:"login"`
	Active    bool            `json:"is_active"`
	TwoFactor bool            `json:"is_2fa_enabled"`
	Name      string          `json:"name"`
	Emails    []UserinfoEmail `json:"emails"`
	Sso       bool            `json:"is_sso_user"`
	Group     string          `json:"default_group_guid"`
}

func GetUserinfo(token string) Userinfo {
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

	var userinfo Userinfo
	if err := json.Unmarshal(bytes, &userinfo); err != nil {
		panic(err)
	}

	return userinfo
}

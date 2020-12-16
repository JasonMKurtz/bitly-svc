package userinfo

/*
	This package queries the /user endpoint to determine the group ID to use for further queries.
*/

import (
	"encoding/json"

	httphelper "../http"
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

// The only public function in this package, this is used to determine the guid of the user in question.
func GetUserinfo(token string) Userinfo {
	var userinfo Userinfo
	if err := json.Unmarshal(httphelper.GetBytes(token, endpoint), &userinfo); err != nil {
		panic(err)
	}

	return userinfo
}

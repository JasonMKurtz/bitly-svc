package bitlinks

/*
   This package provides the data structures and a single public API to retrieve all bitly links
   associated with the owner of the access token.
*/

import (
	"encoding/json"
	"fmt"

	httphelper "../http"
)

type DeepLink struct {
	Id          string `json:"guid"`
	Bitlink     string `json:"bitlink"`
	URI         string `json:"app_uri_path"`
	InstallURL  string `json:"install_url"`
	AppId       string `json:"app_guid"`
	Os          string `json:"os"`
	InstallType string `json:"install_type"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
	BrandId     string `json:"brand_guid"`
}

type Link struct {
	Created     string            `json:"created_at"`
	Id          string            `json:"id"`
	Link        string            `json:"link"`
	CustomLinks []string          `json:"custom_bitlinks"`
	URL         string            `json:"long_url"`
	Title       string            `json:"title"`
	Archived    bool              `json:"archived"`
	CreatedBy   string            `json:"created_by"`
	ClientId    string            `json:"client_id"`
	Tags        []string          `json:"tags"`
	Deeplinks   []DeepLink        `json:"deeplinks"`
	References  map[string]string `json:"references"`
}

type Page struct {
	Prev  string `json:"prev"`
	Next  string `json:"next"`
	Size  int    `json:"size"`
	Page  int    `json:"page"`
	Total int    `json:"total"`
}

type Bitlink struct {
	Links      []Link `json:"links"`
	Pagination Page   `json:"pagination"`
}

// The only function in this package, this returns a list of all the bitly links owned by the user.
func GetBitlinks(groupId, token string) []string {
	endpoint := fmt.Sprintf("https://api-ssl.bitly.com/v4/groups/%s/bitlinks", groupId)
	bytes := httphelper.GetBytes(token, endpoint)

	var bitlink Bitlink
	if err := json.Unmarshal(bytes, &bitlink); err != nil {
		panic(err)
	}

	var ids []string
	for _, link := range bitlink.Links {
		ids = append(ids, link.Id)
	}

	return ids
}

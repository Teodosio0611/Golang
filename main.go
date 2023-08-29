package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Stargazer struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Avatar_url string `json:"avatar_url"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Followers_url string `json:"followers_url"`
	Following_url string `json:"following_url"`
	Gists_url string `json:"gists_url"`
	Starred_url string `json:"starred_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Organizations_url string `json:"organizations_url"`
	Repos_url string `json:"repos_url"`
	Events_url string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type string `json:"type"`
	Site_admin bool `json:"site_admin"`
}

type Stargazers []Stargazer
func main() {
	url := "https://gitee.com/api/v5/repos/tfcolin/ftbt/stargazers?page=1&per_page=20"
	method := "GET"

	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var stargazers Stargazers
	json.Unmarshal(body, &stargazers)
	fmt.Println(stargazers[0].Login)
}


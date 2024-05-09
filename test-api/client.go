package testapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}
type CustomTransport struct {
	r http.RoundTripper
}

func (ct CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-RapidAPI-Key", "a62f144137msh5ce36039416bb23p1dcff1jsnc81c420deeaa")
	req.Header.Add("X-RapidAPI-Host", "fresh-linkedin-profile-data.p.rapidapi.com")
	return ct.r.RoundTrip(req)
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout should be more then 0")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &CustomTransport{
				r: http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetPosts() (PostList, error) {
	postBody := []byte(`{
		"search_keywords": "marketing",
		"sort_by": "Latest",
		"date_posted": "",
		"content_type": "",
		"from_member": "",
		"from_company": "",
		"mentioning_member": "",
		"mentioning_company": 162479,
		"author_company": "",
		"author_industry": "",
		"author_keyword": "",
		"page": 1
	}`)
	resp, err := c.client.Post("https://fresh-linkedin-profile-data.p.rapidapi.com/search-posts", "application/json", bytes.NewBuffer(postBody))

	defer resp.Body.Close()

	if err != nil {
		return PostList{}, errors.New("cant get data")
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return PostList{}, errors.New("cant read body")
	}

	var r PostList

	if err = json.Unmarshal(body, &r); err != nil {
		return PostList{}, errors.New("cant body decode")
	}

	return r, nil
}

func (c Client) GetRecomendations(link string) (RecomendationList, error) {

	url := fmt.Sprintf("https://fresh-linkedin-profile-data.p.rapidapi.com/get-recommendations-given?linkedin_url=%s", link)

	resp, err := c.client.Get(url)

	defer resp.Body.Close()

	if err != nil {
		return RecomendationList{}, errors.New("cant get data")
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return RecomendationList{}, errors.New("cant read body")
	}

	var r RecomendationList

	if err = json.Unmarshal(body, &r); err != nil {
		return RecomendationList{}, errors.New("cant body decode")
	}

	return r, nil
}

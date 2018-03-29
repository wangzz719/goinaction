package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

type (
	data struct {
		Verb        string `json:"verb"`
		Target      target `json:"target"`
		Actor       actor  `json:"actor"`
		ActionText  string `json:"action_text"`
		CreatedTime int64  `json:"created_time"`
		Type        string `json:"type"`
		Id          int64  `json:"id"`
	}

	author struct {
		Headline  string `json:"headline"`
		AvatarUrl string `json:"avatar_url"`
		Name      string `json:"name"`
		Url       string `json:"url"`
		UrlToken  string `json:"url_token"`
		Type      string `json:"type"`
		UserType  string `json:"user_type"`
		Id        string `json:"id"`
	}
	actor struct {
		IsFollowed  bool   `json:"is_followed"`
		Type        string `json:"type"`
		Name        string `json:"name"`
		Headline    string `json:"headline"`
		UrlToken    string `json:"url_token"`
		UserType    string `json:"user_type"`
		Url         string `json:"url"`
		AvatarUrl   string `json:"avatar_url"`
		IsFollowing bool   `json:"is_following"`
		IsOrg       bool   `json:"is_org"`
		Gender      int8   `json:"gender"`
		Id          string `json:"id"`
	}
	target struct {
		Updated           int64       `json:"updated"`
		Description       string      `json:"description"`
		Author            author      `json:"author"`
		Url               string      `json:"url"`
		CommentPermission string      `json:"comment_permission"`
		Title             string      `json:"title"`
		Intro             string      `json:"intro"`
		ImageUrl          string      `json:"image_url"`
		Followers         int64       `json:"followers"`
		Type              string      `json:"type"`
		Id                interface{} `json:"id"`
		ArticlesCount     int32       `json:"articles_count"`
	}
	paging struct {
		IsEnd    bool   `json:"is_end"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	}
	activityResponse struct {
		Paging paging `json:"paging"`
		Data   []data `json:"data"`
	}
)

func main() {
	uri := "https://www.zhihu.com/api/v4/members/wzz/activities?limit=7&after_id=1522067075&desktop=True"
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "oauth c3cef7c66a1843f8b3a9e6a1e3160e20")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	defer resp.Body.Close()
	var activity activityResponse

	// json decoder
	err = json.NewDecoder(resp.Body).Decode(&activity)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(activity)

	resp, err = client.Do(req)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}

	// json string unmarshal
	var activity2 activityResponse
	err = json.Unmarshal(body, &activity2)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(activity2)
}

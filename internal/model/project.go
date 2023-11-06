package model

import "time"

type ESData struct {
	Hits struct {
		Hits []struct {
			Source Project `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type Project struct {
	Slug         string    `json:"slug"`
	Projectname  string    `json:"projectname"`
	Hashtags     []Hashtag `json:"hashtags"`
	Users        []User    `json:"users"`
	Description  string    `json:"description"`
	MaxUpdatedAt time.Time `json:"max_updated_at"`
	Timestamp    time.Time `json:"@timestamp"`
}

type Hashtag struct {
	HashtagName string `json:"hashtagName"`
	HashtagID   int    `json:"hashtagId"`
}

type User struct {
	UserName string `json:"userName"`
	UserID   int    `json:"userId"`
}

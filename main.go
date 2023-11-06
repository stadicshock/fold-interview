package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/elastic/go-elasticsearch/v7"
)

var client *elasticsearch.Client

func init() {
	var err error
	client, err = elasticsearch.NewClient(elasticsearch.Config{
		CloudID: "fold_money_test:ZXVyb3BlLXdlc3QzLmdjcC5jbG91ZC5lcy5pbzo0NDMkNzhjMzhhNGMyZTkwNDllOWFjYmExYzZhNmE3MjFhN2QkMThmNDkzZmYyYWEwNGE1NTgxODE0ZGYxZWY1NjEyNTM=",
		APIKey:  "SUhha3BZc0JGZlRBUzN3LWFTbG46ZVBvTUhnYlRSLVNiMzRJQXNNVTZiUQ==",
	})
	if err != nil {
		log.Println("elasticsearch err:", err)
		panic(err)
	}
	log.Println(client.Info())
}

func main() {

	r := gin.Default()

	r.GET("/projects/created-by/:user_id", SearchProjectsByUser)
	r.GET("/projects/with-hashtags/:hashtags", SearchProjectsByHashtags)
	r.GET("/projects/search", FullTextSearchProjects)

	r.Run(":8080")
}

func SearchProjectsByUser(c *gin.Context) {
	// Handle the search for projects created by a particular user
	user_id := c.Param("user_id")

	query := `{
		"query": {
		  "bool": {
			"must": [
			  {
				"nested": {
				  "path": "users",
				  "query": {
					"term": {
					  "users.userId": "` + user_id + `"
					}
				  }
				}
			  }
			]
		  }
		}
	  }`

	res, err := client.Search(
		client.Search.WithIndex("projects"),
		client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		log.Println("elasticsearch ewew:", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Println("elasticsearch err2:", res)
		return
	}

	var r ESData
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println("elasticsearch:", err)
		return
	}

	hits := r.Hits.Hits

	var projects []Project
	for _, hit := range hits {
		projects = append(projects, hit.Source)
		fmt.Println("hit:", hit)
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})
}

func SearchProjectsByHashtags(c *gin.Context) {
	// Handle the search for projects with multiple hashtags
	hashtags := strings.Split(c.Param("hashtags"), ",")
	query := `{
		"query": {
		  "bool": {
			"must": [
			  {
				"nested": {
				  "path": "hashtags",
				  "query": {
					"term": {
					  "hashtags.hashtagName": "` + hashtags[0] + `"
					}
				  }
				}
			  }
			]
		  }
		}
	  }`

	res, err := client.Search(
		client.Search.WithIndex("projects"),
		client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		log.Println("elasticsearch ewew:", err)
		return
	}
	defer res.Body.Close()

	// Check the response and process the data
	if res.IsError() {
		log.Println("elasticsearch err2:", res)
		return
	}

	// Decode the response body, which contains the search results
	var r ESData
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println("elasticsearch:", err)
		return
	}

	hits := r.Hits.Hits

	var projects []Project
	for _, hit := range hits {
		projects = append(projects, hit.Source)
		fmt.Println("hit:", hit)
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})
}

func FullTextSearchProjects(c *gin.Context) {
	// Handle the full-text fuzzy search for projects
	searchQuery := c.Query("q")

	query := `{
		"query": {
		  "multi_match": {
			"query": "` + searchQuery + `",
			"fields": ["slug^2", "description"],
			"fuzziness": "AUTO"
		  }
		}
	  }`

	res, err := client.Search(
		client.Search.WithIndex("projects"),
		client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		log.Println("elasticsearch ewew:", err)
		return
	}
	defer res.Body.Close()

	// Check the response and process the data
	if res.IsError() {
		log.Println("elasticsearch err2:", res)
		return
	}

	// Decode the response body, which contains the search results
	var r ESData
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println("elasticsearch:", err)
		return
	}

	hits := r.Hits.Hits

	var projects []Project
	for _, hit := range hits {
		projects = append(projects, hit.Source)
		fmt.Println("hit:", hit)
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})
}

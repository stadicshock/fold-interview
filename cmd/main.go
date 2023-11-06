package main

import (
	"fold/internal/config"
	server "fold/internal/http"
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {

	esClient, err := getESClient()
	if err != nil {
		panic(err)
	}

	s, err := server.New(esClient)
	if err != nil {
		panic(err)
	}
	s.Start()
}

func getESClient() (*elasticsearch.Client, error) {

	client, err := elasticsearch.NewClient(elasticsearch.Config{
		CloudID: config.ESCloudID,
		APIKey:  config.ESApiKey,
	})
	if err != nil {
		log.Println("Elasticsearch NewClient err:", err)
		return nil, err
	}

	_, err = client.Ping()
	if err != nil {
		log.Println("Elasticsearch Ping err:", err)
		return nil, err
	}
	return client, nil
}

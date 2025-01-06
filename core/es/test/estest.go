package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://129.3.212.60:9200/",
		},
		Username: "admin",
		Password: "",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Check Elasticsearch connection
	info, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer info.Body.Close()

	var mapResp map[string]interface{}
	if err := json.NewDecoder(info.Body).Decode(&mapResp); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("Elasticsearch Info: %s", mapResp)

	// Define the query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"wildcard": map[string]interface{}{
				"username": map[string]interface{}{
					"value": "*bin*",
				},
			},
		},
	}

	// Encode the query to JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("win_user"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the error response body: %s", err)
		}
		log.Fatalf("Elasticsearch search error: %s", e["error"])
	}

	// Parse and print search results
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("Elasticsearch search result: %s", result)
}

package main

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	// GET запрос
	resp, err := client.R().
		SetQueryParam("userId", "1").
		Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalf("Error on GET request: %v", err)
	}
	fmt.Println("GET Response Info:")
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Body:", resp.String())

	// POST запрос
	post := map[string]interface{}{
		"userId": 1,
		"title":  "foo",
		"body":   "bar",
	}
	resp, err = client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(post).
		Post("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalf("Error on POST request: %v", err)
	}
	fmt.Println("POST Response Info:")
	fmt.Println("Status Code:", resp.StatusCode())
	fmt.Println("Body:", resp.String())
}

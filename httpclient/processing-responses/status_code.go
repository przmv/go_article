package main

import (
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	httpClient := NewHTTPClient()

	// GET-запрос
	response, err := httpClient.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received non-200 response code: %d\n", response.StatusCode)
		return
	}

	fmt.Printf("Received a successful response. Status code: %d\n", response.StatusCode)
}

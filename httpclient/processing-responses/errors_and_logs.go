package main

import (
	"fmt"
	"log"
	"os"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	if err := run(); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	client := NewHTTPClient()

	post, err := client.GetBlogPost(1)
	if err != nil {
		return fmt.Errorf("error occurred while getting post: %w", err)
	}

	fmt.Printf("ID: %d\nUser ID: %d\nTitle: %s\nBody: %s\n", post.ID, post.UserID, post.Title, post.Body)

	return nil
}

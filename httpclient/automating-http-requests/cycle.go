package main

import (
	"fmt"
	"log"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	client := NewHTTPClient()

	for i := 1; i <= 5; i++ {
		post, err := client.GetBlogPost(i)
		if err != nil {
			log.Printf("Error getting post %d: %v", i, err)
			continue
		}

		fmt.Printf("Request to post %d returned:\nID: %d \n%s \n\n",
			i, post.ID, post.Title)
	}
}

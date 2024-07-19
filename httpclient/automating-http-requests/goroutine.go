package main

import (
	"fmt"
	"log"
	"sync"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

// fetchPost обрабатывает получение поста через метод GetBlogPost и выводит результат.
func fetchPost(client *HTTPClient, postID int, wg *sync.WaitGroup) {
	defer wg.Done()

	post, err := client.GetBlogPost(postID)
	if err != nil {
		log.Printf("Error getting post %d: %v", postID, err)
		return
	}

	fmt.Printf("Request to post %d returned:\nID: %d\nUser ID: %d\nTitle: %s\nBody: %s\n\n",
		postID, post.ID, post.UserID, post.Title, post.Body)
}

func main() {
	client := NewHTTPClient()
	var wg sync.WaitGroup

	postIDs := []int{1, 2, 3, 4, 5}

	for _, postID := range postIDs {
		wg.Add(1)
		go fetchPost(client, postID, &wg)
	}

	wg.Wait()
}

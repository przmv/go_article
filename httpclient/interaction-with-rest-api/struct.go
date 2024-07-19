package main

import (
	"fmt"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	// Создаем HTTP клиент
	client := NewHTTPClient()

	// Получаем данные поста
	post, err := client.GetBlogPost(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Выводим данные поста
	fmt.Printf("Post ID: %d\n", post.ID)
	fmt.Printf("User ID: %d\n", post.UserID)
	fmt.Printf("Title: %s\n", post.Title)
	fmt.Printf("Body: %s\n", post.Body)
}

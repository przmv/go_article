package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"httpclinet/client"
)

func main() {
	// Инициализация пользовательского HTTP-клиента
	httpClient := client.NewHTTPClient(&http.Client{
		Timeout: 10 * time.Second,
	})

	ctx := context.Background()

	input := &client.BlogPost{
		Title:  "foo",
		Body:   "bar",
		UserID: 1,
	}

	// Создание нового ресурсв с использованием пользовательского HTTP-клиента
	blogPost, _, err := httpClient.CreateBlogPost(ctx, input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Created Blog Post:")
	fmt.Printf("  ID: %d\n", blogPost.ID)
	fmt.Printf("  Title: %s\n", blogPost.Title)
	fmt.Printf("  Body: %s\n", blogPost.Body)
	fmt.Printf("  User ID: %d\n", blogPost.UserID)
}

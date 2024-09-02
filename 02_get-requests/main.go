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

	// Получение существующего поста с использованием пользовательского HTTP-клиента
	blogPost, _, err := httpClient.GetBlogPost(ctx, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Blog Post:")
	fmt.Printf("  ID: %d\n", blogPost.ID)
	fmt.Printf("  Title: %s\n", blogPost.Title)
	fmt.Printf("  Body: %s\n", blogPost.Body)
	fmt.Printf("  User ID: %d\n", blogPost.UserID)

	// Получение несуществующего поста с использованием пользовательского HTTP-клиента
	blogPost, _, err = httpClient.GetBlogPost(ctx, -1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Blog Post:", blogPost)
}

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
	body, err := httpClient.GetBlogPost(ctx, 1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", body)

	// Получение несуществующего поста с использованием пользовательского HTTP-клиента
	body, err = httpClient.GetBlogPost(ctx, -1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", body)
}

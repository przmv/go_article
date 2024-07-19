package main

import (
	"fmt"
)

func main() {
	// Инициализация пользовательского HTTP-клиента
	httpClient := NewHTTPClient()

	// Получение поста с использованием пользовательского HTTP-клиента
	body, err := httpClient.GetBlogPost(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", body)
}

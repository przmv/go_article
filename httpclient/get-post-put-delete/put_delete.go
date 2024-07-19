package main

import (
	"fmt"
)

func main() {
	client := NewHTTPClient()

	// Пример PUT-запроса
	jsonToPut := []byte(`{"id": 1, "title": "foo", "body": "bar", "userId": 1}`)
	putResp, err := client.PutJSON("https://jsonplaceholder.typicode.com/posts/1", jsonToPut)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("PUT Response:", putResp)
		fmt.Println("PUT Response:", putResp)
	}

	// Пример DELETE-запроса
	deleteResp, err := client.Delete("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("DELETE Response:", deleteResp)
	}
}

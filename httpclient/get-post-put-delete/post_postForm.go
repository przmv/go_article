package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	httpClient := NewHTTPClient()

	postData := map[string]interface{}{
		"title":  "foo",
		"body":   "bar",
		"userId": 1,
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	responsePost, err := httpClient.PostJSON("https://jsonplaceholder.typicode.com/posts", jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", responsePost)

	formData := map[string]string{
		"title":  "foo",
		"body":   "bar",
		"userId": "1",
	}

	responsePostForm, err := httpClient.PostForm("https://jsonplaceholder.typicode.com/posts", formData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response Body:", responsePostForm)
}

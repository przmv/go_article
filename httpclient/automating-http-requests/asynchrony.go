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

type Result struct {
	PostID int
	Post   *Post
	Err    error
}

// fetchPost обрабатывает получение поста через метод GetBlogPost и отправляет результат в канал.
func fetchPost(client *HTTPClient, postID int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	post, err := client.GetBlogPost(postID)
	results <- Result{PostID: postID, Post: post, Err: err}
}

func main() {
	client := NewHTTPClient()
	var wg sync.WaitGroup

	postIDs := []int{1, 2, 3, 4, 5}
	results := make(chan Result, len(postIDs))

	// Запуск горутин для параллельного выполнения запросов
	for _, postID := range postIDs {
		wg.Add(1)
		go fetchPost(client, postID, results, &wg)
	}

	// Функция для закрытия канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Обработка результатов по мере их поступления
	for result := range results {
		if result.Err != nil {
			log.Printf("Error fetching post %d: %v\n", result.PostID, result.Err)
			continue
		}
		fmt.Printf("Request to post %d returned:\nID: %d\nUser ID: %d\nTitle: %s\nBody: %s\n\n",
			result.PostID, result.Post.ID, result.Post.UserID, result.Post.Title, result.Post.Body)
	}
}

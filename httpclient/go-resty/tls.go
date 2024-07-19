package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func main() {
	// Создание клиента с настроенным TLS
	client := resty.New()

	// Настройка транспортного уровня безопасности
	client.SetTransport(&http.Transport{
		// Использование стандартного конфигурации TLS
		TLSClientConfig: &tls.Config{
			// Дополнительные параметры конфигурации можно установить здесь
			MinVersion: tls.VersionTLS12, // Пример: минимальная версия TLS 1.2
		},
	})

	token := "your_auth_token_here"

	// Отправка GET запроса с обработкой ошибок и проверкой TLS
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if resp.StatusCode() != http.StatusOK {
		log.Fatalf("Non-200 response: %d", resp.StatusCode())
	}

	// Обработка тела ответа
	fmt.Printf("Response: %s\n", resp.String())
}

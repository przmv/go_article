package main

import (
	"encoding/xml"
	"fmt"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Response struct {
	XMLName xml.Name `xml:"objects"`
	Objects []Object `xml:"object"`
}

type Object struct {
	ID        int    `xml:"id"`
	Name      string `xml:"name"`
	Email     string `xml:"email"`
	Avatar    string `xml:"avatar"`
	CreatedAt string `xml:"created-at"`
	UpdatedAt string `xml:"updated-at"`
}

func main() {
	httpClient := NewHTTPClient()

	var response Response

	err := httpClient.GetXML("https://thetestrequest.com/authors.xml", &response)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, obj := range response.Objects {
		fmt.Printf("ID: %d, Name: %s, Email: %s, Avatar: %s, CreatedAt: %s, UpdatedAt: %s\n",
			obj.ID, obj.Name, obj.Email, obj.Avatar, obj.CreatedAt, obj.UpdatedAt)
	}
}

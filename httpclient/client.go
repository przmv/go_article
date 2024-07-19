package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type HTTPClient struct {
	Client *http.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *HTTPClient) GetBlogPost(postID int) (string, error) {
	resp, err := c.Client.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID))
	if err != nil {
		return "", fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

//Измененный GetPostBlog для структуры
// func (c *HTTPClient) GetBlogPost(postID int) (*Post, error) {
// 	resp, err := c.Client.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", postID))
// 	if err != nil {
// 		return nil, fmt.Errorf("error making GET request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	var post Post
// 	err = json.NewDecoder(resp.Body).Decode(&post)
// 	if err != nil {
// 		return nil, fmt.Errorf("error decoding response body: %w", err)
// 	}

// 	return &post, nil
// }

func (c *HTTPClient) PostJSON(url string, jsonData []byte) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating POST request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making POST request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

func (c *HTTPClient) PostForm(myUrl string, formData map[string]string) (string, error) {
	form := url.Values{}
	for key, value := range formData {
		form.Set(key, value)
	}

	resp, err := c.Client.PostForm(myUrl, form)
	if err != nil {
		return "", fmt.Errorf("error making POST form request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

func (c *HTTPClient) PutJSON(myUrl string, jsonData []byte) (string, error) {
	req, err := http.NewRequest("PUT", myUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating PUT request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making PUT request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

func (c *HTTPClient) Delete(myUrl string) (string, error) {
	req, err := http.NewRequest("DELETE", myUrl, nil)
	if err != nil {
		return "", fmt.Errorf("error creating DELETE request: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making DELETE request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}

func (c *HTTPClient) Get(url string) (*http.Response, error) {
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %w", err)
	}
	return resp, nil
}

func (c *HTTPClient) GetXML(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating GET request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	resp, err := c.Client.Do(req)
	if err != nil {
		return fmt.Errorf("error making GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = xml.Unmarshal(body, v)
	if err != nil {
		return fmt.Errorf("error unmarshalling XML response: %w", err)
	}

	return nil
}

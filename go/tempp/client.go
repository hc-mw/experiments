package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	pp "github.com/k0kubun/pp"
)

func NewClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

func MakeRequest() error {
	client := NewClient(30 * time.Second)

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "Application/json")

	response, err := client.Do(req)
	if err != nil {
		return err
	}

	pp.Println(response)

	var b bytes.Buffer

	io.Copy(&b, response.Body)

	fmt.Println(b.String())

	return nil
}

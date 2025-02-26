package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	requestURL := "http://localhost:3333"
	salute := "Hello Client"
	// jsonBody := []byte(`{"client_message": "hello, server!"}`)
	jsonBody := []byte(fmt.Sprintf(`{"client_message": "%s"}`, salute))
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making request: %s\n", err)
		os.Exit(1)
	}

	body, _ := io.ReadAll(res.Body)
	fmt.Printf("client: response body: %s\n", body)
}

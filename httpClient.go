package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	requestURL := "http://localhost:3333"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making request: %s\n", err)
		os.Exit(1)
	}

	body, _ := io.ReadAll(res.Body)
	fmt.Printf("client: response body: %s\n", body)
}

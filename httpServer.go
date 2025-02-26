package main

import (
	"fmt"
	"io"
	"net/http"
)

const serverPort = 3333

func main() {
	// Iniciar el servidor en un goroutine
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("server: %s /\n", r.Method)
			fmt.Printf("server: query id: %s\n", r.URL.Query().Get("id"))
			fmt.Printf("server: content-type: %s\n", r.Header.Get("content-type"))

			body, _ := io.ReadAll(r.Body)
			fmt.Printf("server: request body: %s\n", body)

			fmt.Fprintf(w, `{"message": "hello!"}`)
		})
		http.ListenAndServe(fmt.Sprintf(":%d", serverPort), mux)
	}()

	select {} // Mantiene el servidor ejecutándose
}

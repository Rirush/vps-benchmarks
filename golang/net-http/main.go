package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello, World!"))
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}

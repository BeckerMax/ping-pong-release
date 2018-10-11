package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	ping()
}

func ping() {
	pongIP := ""
	pingReader := strings.NewReader("Hello, Reader!")
	resp, err := http.Post(pongIP, "text/plain", pingReader)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(resp)
}

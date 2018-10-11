package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	doAll10Seconds(ping)
}

func doAll10Seconds(repeatMe func()) {
	for {
		repeatMe()
		time.Sleep(10 * time.Second)
	}
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

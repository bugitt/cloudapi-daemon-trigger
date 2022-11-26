package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func trigger(url string) {
	// do post request and recover
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic: ", r)
		}
	}()
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		log.Println("Error when do post request: ", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error when read the response: ", err)
	}
	log.Println("Response: ", string(body))
}

func main() {
	if len(os.Args) != 3 {
		panic("Usage: ./cloudapi-daemon-trigger <url> <interval>")
	}
	url, interval := os.Args[1], os.Args[2]
	intervalInt, err := strconv.ParseFloat(interval, 64)
	if err != nil {
		panic("Interval must be a number")
	}
	ticker := time.NewTicker(time.Duration(intervalInt*1000) * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				trigger(url)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	<-quit
}

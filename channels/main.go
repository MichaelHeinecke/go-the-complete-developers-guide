package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://www.golang.org",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.facebook.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(url, c)
	}

	// Infinite loop, runs everytime a value is received from the channel.
	for url := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkUrl(link, c)
		}(url)
	}
}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // -> Blocking call.

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep its up"
}

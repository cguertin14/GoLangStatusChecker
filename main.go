package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string) // Création du channel

	for _, link := range links {
		go checkLink(link, c) // Création d'une go routine
	}

	for l := range c {
		go func(link string) { // Création d'une go routine
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // -> Blocking call.

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}


package main

import "fmt"

func main() {

    // Création d'un channel
    messages := make(chan string)

	// 1- Création d'une go routine 
	// 2- Envoi d'une valeur au channel "messages"
    go func() { messages <- "ping" }()

    // Réception du message provenant du channel
	msg := <-messages
	
	// Impression de ce message
    fmt.Println(msg)
}
package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	// for {
	// 	go checkLink(<-c, c)
	// }
	// This is same as above but more explicit and readable
	// for l := range c {
	// 	go checkLink(l, c)
	// }
	for l := range c {
		//anon func called functiona literal in GO
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l) // l from range chancel is being passed here
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down", err)
		c <- link
		return
	}
	fmt.Println(link, "is up and running ")
	c <- link
}

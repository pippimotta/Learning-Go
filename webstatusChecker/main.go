package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		//to make it continue to execute
		// time.Sleep(5 * time.Second) if the pause is made within the main func, the main routine will not be able to receive it
		//use a lambda function to make the pause
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //pass link as a param to avoid directly access the same variable from a different child routine; only with channel or params

	}

}

func checkLink(link string, c chan string) {
	//time.Sleep(5 * time.Second) //keep the main routine always running
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link

}

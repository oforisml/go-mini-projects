package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://marca.com",
		"http://realmadrid.com",
		"http://bing.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://yahoo.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkStatus(link, c)

	}
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkStatus(link, c)
		}(l)

	}

}

func checkStatus(link string, c chan string) {
	_, e := http.Get(link)
	if e != nil {
		fmt.Println("The link (", link, ") is Down")
		fmt.Println("ERROR: ", e)
		c <- link
		return
	}
	fmt.Println("Link: ", link, "is UP")
	c <- link
}

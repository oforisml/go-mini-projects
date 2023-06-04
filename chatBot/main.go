package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

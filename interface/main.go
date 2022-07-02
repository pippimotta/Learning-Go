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

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	//custom logic for generating an spanish greeting
	return "Hola!"
}

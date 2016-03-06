package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "ping" }()
	go func() { messages <- "pong" }()
	go func() { messages <- "yabadabadoo" }()
	go func() { messages <- "kuristsu" }()

	m := <-messages
	fmt.Println(m)
	m = <-messages
	fmt.Println(m)
	m = <-messages
	fmt.Println(m)
	m = <-messages
	fmt.Println(m)
}

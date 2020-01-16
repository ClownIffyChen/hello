package main

import "fmt"

const englishisHelloPrefix = "Hello, "

// hello, world
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishisHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}

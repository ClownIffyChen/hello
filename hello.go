package main

import "fmt"

// hello
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}

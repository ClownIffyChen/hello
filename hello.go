package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	spanishHelloPrefix = "Hola, "
	helloPrefix        = "Hello, "
	frenchHelloPrefix  = "French, "
)

// hello, world
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := helloPrefix
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}

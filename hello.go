package main

import "fmt"

const helloPrefix = "Hello, "
const defaultHello = helloPrefix + "World"

func Hello(name string) string {

	if name == "" {
		return defaultHello
	}

	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("World"))
}

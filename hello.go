package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const defaultHello = englishPrefix + "World"
const defaultName = "World"

func Hello(name string, language string) string {

	if len(name) <= 0 {
		name = defaultName
	}

	if language == "es_ES" {
		return spanishPrefix + name
	}

	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}

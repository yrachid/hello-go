package main

import "fmt"

const frenchPrefix = "Bonjour, "
const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const defaultHello = englishPrefix + "World"
const defaultName = "World"

const spanish = "es_ES"
const french = "fr_FR"

func Hello(name string, language string) string {

	if len(name) <= 0 {
		name = defaultName
	}

	if language == spanish {
		return spanishPrefix + name
	}

	if language == french {
		return frenchPrefix + name
	}

	return englishPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}

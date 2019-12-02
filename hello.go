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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (greeting string) {
	switch language {
	case french:
		greeting = frenchPrefix
	case spanish:
		greeting = spanishPrefix
	default:
		greeting = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}

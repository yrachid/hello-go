package main

import "testing"

var isEqualTo = func(value string) string {
	return value
}

var assertThat = func(t *testing.T, actual, expected string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Got %q expected %q", actual, expected)
	}
}

func TestHello(t *testing.T) {

	t.Run("Hello, Doot when Hello('Doot', '')", func(t *testing.T) {
		assertThat(t, Hello("Doot", ""), isEqualTo("Hello, Doot"))
	})

	t.Run("Hello, World when Hello('', '')", func(t *testing.T) {
		assertThat(t, Hello("", ""), isEqualTo("Hello, World"))
	})

	t.Run("Hola, Mundo", func(t *testing.T) {
		assertThat(t, Hello("Mundo", "es_ES"), isEqualTo("Hola, Mundo"))
	})

	t.Run("Bonjour, Monde", func(t *testing.T) {
		assertThat(t, Hello("Monde", "fr_FR"), isEqualTo("Bonjour, Monde"))
	})

}

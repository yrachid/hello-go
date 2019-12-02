package main

import "testing"

func TestHello(t *testing.T) {

	isEqualTo := func(value string) string {
		return value
	}

	assertThat := func(t *testing.T, actual, expected string) {
		t.Helper()

		if actual != expected {
			t.Errorf("Got %q expected %q", actual, expected)
		}
	}

	t.Run("Hello, Doot when Hello('Doot')", func(t *testing.T) {
		assertThat(t, Hello("Doot"), isEqualTo("Hello, Doot"))
	})

	t.Run("Hello, World when Hello('')", func(t *testing.T) {
		assertThat(t, Hello(""), isEqualTo("Hello, World"))
	})

}

package hello

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say my name!", func(t *testing.T) {
		t.Logf("Hello JP")
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)

	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		got := Hello("JP", "spanish")
		want := "Hola JP"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in french", func(t *testing.T) {
		got := Hello("JP", "french")
		want := "Bonjour JP"

		assertCorrectMessage(t, got, want)
	})

}

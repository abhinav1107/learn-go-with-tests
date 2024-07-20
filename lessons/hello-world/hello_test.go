package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Run Without Arguments", func(t *testing.T) {
		got := Hello(Params{})
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Run with Just name", func(t *testing.T) {
		got := Hello(Params{name: "abhinav"})
		want := "Hello, Abhinav"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Run to greet world in French", func(t *testing.T) {
		got := Hello(Params{greetLang: "French"})
		want := "Bonjour, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Run to greet Jim in Spanish", func(t *testing.T) {
		got := Hello(Params{name: "jim", greetLang: "Spanish"})
		want := "Hola, Jim"
		assertCorrectMessage(t, got, want)
	})
}

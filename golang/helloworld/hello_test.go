package main

import "testing"

// need to be in a file with the name any_test.go
// func name must start with Test
// takes only one args

func TestHello(t *testing.T) {
	// it's possible to have multiple tests, AKA subtest inside one function
	t.Run("without args", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello There!"
		assertString(t, got, want)
	})
	t.Run("with args", func(t *testing.T) {
		got := Hello("Bro", "")
		want := "Hello There! Bro!"
		assertString(t, got, want)
	})
	t.Run("with lang", func(t *testing.T) {
		got := Hello("Bro", "Bahasa")
		want := "Hai Kamu! Bro!"
		assertString(t, got, want)
	})
	subTest(t, "with lange EN", Hello("Bro", "English"), "Hello There! Bro!")
}

// some kind of wrapper function of t.Run, kinda useless tho
func subTest(t *testing.T, name, got, want string) {
	t.Run(name, func(t *testing.T) {
		assertString(t, got, want)
	})
}

// doesn't care about the order of function calls
// var type can be defined for two var
func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}

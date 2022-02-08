package main

import "testing"

func TestApp(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run(
		"Test a valid payload", func(t *testing.T) {
			got := App("test.txt", "This test should pass", "Bearer auth-token-123")
			want := "Success"
			assertCorrectMessage(t, got, want)
		})
	t.Run(
		"Test with missing filename", func(t *testing.T) {
			got := App("", "I am an error", "Bearer auth-token-123")
			want := "Payload is invalid."
			assertCorrectMessage(t, got, want)
		})
	t.Run(
		"Test with missing content", func(t *testing.T) {
			got := App("test.txt", "", "Bearer auth-token-123")
			want := "Payload is invalid."
			assertCorrectMessage(t, got, want)
		})
	t.Run(
		"Test with wrong token", func(t *testing.T) {
			got := App("test.txt", "I am an error", "wrong-token")
			want := "Payload is invalid."
			assertCorrectMessage(t, got, want)
		})
}

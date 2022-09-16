package config

import "testing"

func TestVersion(t *testing.T) {
	assertCorrect := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("is ok version", func(t *testing.T) {
		got := ShowVersion()
		want := "development"
		assertCorrect(t, got, want)
	})

	// t.Run("is not ok version", func(t *testing.T) {
	// 	got := ShowVersion()
	// 	want := "not development"
	// 	assertCorrect(t, got, want)
	// })
}

func TestVersion2(t *testing.T) {
	got := ShowVersion()
	want := "development"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

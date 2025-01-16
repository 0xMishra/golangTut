package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"
		assertSrtings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		if err == nil {
			t.Fatal("expected error, but didn't get one")
		}

		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "test"
		def := "this is just a test"
		err := dict.Add(word, def)

		assertErrors(t, err, nil)
		assertDefinition(t, dict, word, def)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}
		err := dict.Add(word, def)
		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dict, word, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}

		newDef := "some test"

		err := dict.Update(word, newDef)
		assertErrors(t, err, nil)
		assertDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		newDef := "some test"

		err := dict.Update(word, newDef)
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is just a test"
		dict := Dictionary{word: def}

		err := dict.Delete(word)
		assertErrors(t, err, nil)
		assertDelete(t, dict, word)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dict := Dictionary{}

		err := dict.Delete(word)
		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func assertDelete(t testing.TB, dict Dictionary, word string) {
	t.Helper()
	_, err := dict.Search(word)
	if err == nil {
		t.Fatal("expected not found error, but didn't get one")
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word, def string) {
	t.Helper()
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("expected no errors, but got one")
	}

	assertSrtings(t, got, def)
}

func assertSrtings(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func assertErrors(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got.Error(), want.Error())
	}
}

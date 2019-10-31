package dir

import "testing"

func TestShow(t *testing.T) {
	if err := Read(".."); err != nil {
		t.Fatal(err)
	}
}

func TestShowRecursively(t *testing.T) {
	if err := ReadRecursively(".."); err != nil {
		t.Fatal(err)
	}
}

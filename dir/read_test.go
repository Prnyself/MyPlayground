package dir

import "testing"

func TestRead(t *testing.T) {
	if err := Read(".."); err != nil {
		t.Fatal(err)
	}
}

func TestReadRecursivelyRecursively(t *testing.T) {
	if err := ReadRecursively(".."); err != nil {
		t.Fatal(err)
	}
}

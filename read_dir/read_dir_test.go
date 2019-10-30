package read_dir

import "testing"

func TestShow(t *testing.T) {
	if err := Show(".."); err != nil {
		t.Fatal(err)
	}
}

func TestShowRecursively(t *testing.T) {
	if err := ShowRecursively(".."); err != nil {
		t.Fatal(err)
	}
}

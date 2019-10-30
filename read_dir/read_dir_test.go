package read_dir

import "testing"

func TestShow(t *testing.T) {
	if err := Show(".."); err != nil {
		t.Fatal(err)
	}
}

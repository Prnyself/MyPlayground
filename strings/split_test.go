package strings

import "testing"

func TestSplit(t *testing.T) {
	if err := Split("127.0.0.1", ":"); err != nil {
		t.Fatal(err)
	}
}

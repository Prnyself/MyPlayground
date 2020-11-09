package strings

import "testing"

func TestRangeString(t *testing.T) {
	str := "x-qs-meta-data-custom"
	for i, b := range str {
		t.Logf("i: %d, b: %c", i, b)
	}
}

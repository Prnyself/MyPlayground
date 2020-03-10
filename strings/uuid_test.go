package strings

import "testing"

func TestParseID(t *testing.T) {
	cases := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range cases {
		if _, err := ParseID(); err != nil {
			t.Fatal(err, tt.name)
		}
	}
}

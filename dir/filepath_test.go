package dir

import "testing"

func TestPath(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"."},
		{"/"},
		{"filepath.go"},
	}

	for _, tt := range cases {
		if err := Path(tt.input); err != nil {
			t.Fatal(err)
		}
	}
}

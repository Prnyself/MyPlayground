package size

import "testing"

func TestReadableSize(t *testing.T) {
	cases := []struct {
		input int64
	}{
		{1073741824},
	}

	for _, tt := range cases {
		if err := ReadableSize(tt.input); err != nil {
			t.Fatal(err)
		}
	}
}

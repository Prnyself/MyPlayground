package os

import "testing"

func TestStat(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"stat.go"},
		{"/Users/lance/Desktop/QingCloud/projects/modgo/qsctl/move_test/1/out.txt"},
	}

	for _, tt := range cases {
		if err := Stat(tt.input); err != nil {
			t.Fatal(err)
		}
	}
}

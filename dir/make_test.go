package dir

import (
	"testing"
)

func TestMake(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"/Users/lance/Desktop/QingCloud/projects/modgo/MyPlayground/test1/test12/abc.go"},
		{"/Users/lance/Desktop/QingCloud/projects/modgo/MyPlayground/test2/test22/"},
	}

	for _, tt := range cases {
		if err := Make(tt.input); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCreateFile(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"/Users/lance/Desktop/QingCloud/projects/modgo/MyPlayground/test1/test12/abc.go"},
	}

	for _, tt := range cases {
		if err := CreateFile(tt.input); err != nil {
			t.Fatal(err)
		}
	}
}

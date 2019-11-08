package strings

import "testing"

func TestTrimPrefix(t *testing.T) {
	cases := []struct {
		s      string
		prefix string
	}{
		{s: "abc", prefix: "/"},
		{s: "/abc", prefix: "/"},
		{s: "////abc", prefix: "/"},
	}

	for _, tt := range cases {
		if err := TrimPrefix(tt.s, tt.prefix); err != nil {
			t.Fatal(err)
		}
	}
}

func TestTrimLeft(t *testing.T) {
	cases := []struct {
		s string
		l string
	}{
		{s: "abc", l: "/"},
		{s: "/abc", l: "/"},
		{s: "/abc/def", l: "/"},
		{s: "///abc/def", l: "/"},
	}
	for _, tt := range cases {
		if err := TrimLeft(tt.s, tt.l); err != nil {
			t.Fatal(err)
		}
	}
}

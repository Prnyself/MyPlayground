package strings

import "testing"

func TestSlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"no truncate", args{s: "abc"}, "abc"},
		{"len 10", args{s: "abcdefghij"}, "abcdefghij"},
		{"longer than 10", args{s: "abcdefghijklmn"}, "...efghijklmn"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Slice(tt.args.s); got != tt.want {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

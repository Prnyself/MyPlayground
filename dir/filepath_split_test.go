package dir

import (
	"os"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		input string
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal("get work dir failed", err)
	}
	t.Log(wd)
	tests := []struct {
		name     string
		args     args
		wantWd   string
		wantPath string
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:     ".",
			args:     args{input: "."},
			wantWd:   wd + "/",
			wantPath: "",
			wantErr:  false,
		},
		{
			name:     "/abs/path/file.go",
			args:     args{input: "/abs/path/file.go"},
			wantWd:   "/abs/path/",
			wantPath: "file.go",
			wantErr:  false,
		},
		{
			name:     "rel/path/file.go",
			args:     args{input: "rel/path/file.go"},
			wantWd:   wd + "/rel/path/",
			wantPath: "file.go",
			wantErr:  false,
		},
		{
			name:     "/abs/path/",
			args:     args{input: "/abs/path/"},
			wantWd:   "/abs/path/",
			wantPath: "",
			wantErr:  false,
		},
		{
			name:     "rel/path/",
			args:     args{input: "rel/path/"},
			wantWd:   wd + "/rel/path/",
			wantPath: "",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWd, gotPath, err := Split(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWd != tt.wantWd {
				t.Errorf("Split() gotWd = %v, want %v", gotWd, tt.wantWd)
			}
			if gotPath != tt.wantPath {
				t.Errorf("Split() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

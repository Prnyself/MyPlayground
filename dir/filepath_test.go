package dir

import "testing"

func TestPath(t *testing.T) {
	cases := []struct {
		input string
	}{
		{"."},
		{"/"},
		{"filepath.go"},
		{"/abs/path/to/filepath.go"},
		{"rel/path/to/filepath.go"},
		{"rel/path/to/path/"},
		{"/abs/path/to/path/"},
	}

	for _, tt := range cases {
		if err := Path(tt.input); err != nil {
			t.Fatal(err)
		}
	}
}

func TestAbsAndDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "rel/path/to/file.go",
			args:    args{path: "rel/path/to/file.go"},
			wantErr: false,
		},
		{
			name:    "rel/path/",
			args:    args{path: "rel/path/"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AbsAndDir(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("AbsAndDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

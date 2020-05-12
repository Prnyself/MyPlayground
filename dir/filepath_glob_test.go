package dir

import (
	"reflect"
	"testing"
)

func TestGlob(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "one *",
			args:    args{"./filepath_*"},
			want:    []string{"filepath_glob.go", "filepath_glob_test.go", "filepath_split.go", "filepath_split_test.go", "filepath_test.go"},
			wantErr: false,
		},
		{
			name:    "two *",
			args:    args{"../f*/*a*.go"},
			want:    []string{"../flags/flag.go", "../func/func_param.go", "../func/func_param_test.go"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Glob(tt.args.path)
			t.Log(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Glob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Glob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package strings

import "testing"

func TestConvToFloat(t *testing.T) {
	type args struct {
		s string
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "normal",
			args:    args{"6.824", 64},
			want:    6.824,
			wantErr: false,
		},
		{
			name:    "broken",
			args:    args{"6.824", 32},
			want:    6.823999881744385,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvToFloat(tt.args.s, tt.args.b)
			t.Log("got:", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvToFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvToFloat() got = %v, want %v", got, tt.want)
			}
		})
	}
}

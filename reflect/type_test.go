package reflect

import "testing"

func TestShowTypeOf(t *testing.T) {
	cases := []struct {
		input interface{}
	}{
		{[]string{"1"}},
	}
	for _, tt := range cases {
		if err := ShowTypeOf(tt.input); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSliceToInterface(t *testing.T) {
	cases := []struct {
		input interface{}
	}{
		{[]string{"1", "2"}},
	}
	for _, tt := range cases {
		output, err := SliceToInterface(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(len(output))
	}
}

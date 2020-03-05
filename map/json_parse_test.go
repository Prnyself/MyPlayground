package _map

import "testing"

func TestUnmarshal(t *testing.T) {
	cases := []struct {
		input map[string]interface{}
	}{
		{
			input: map[string]interface{}{
				"name": "abc",
				"type": map[string]interface{}{"aaa": 123},
			},
		},
	}
	for _, tt := range cases {
		if err := Unmarshal(tt.input); err != nil {
			t.Fatal(err)
		}

	}
}

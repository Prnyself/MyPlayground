package _map

import (
	"testing"
)

func TestInitActions(t *testing.T) {
	if err := InitActions(); err != nil {
		t.Fatal(err)
	}
	t.Log(ActionMap)
}

func TestConRead(t *testing.T) {
	if err := InitActions(); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 4; i++ {
		// go func() {
		for key, value := range ActionMap {
			t.Logf("k: %s, v: %v", key, value)
		}
		// }()
	}
}

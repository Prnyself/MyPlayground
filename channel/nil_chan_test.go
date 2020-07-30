package channel

import "testing"

func TestShowNilChan(t *testing.T) {
	if err := ShowNilChan(); err != nil {
		t.Fatal(err)
	}
}

func TestShowNilWg(t *testing.T) {
	if err := ShowNilWg(); err != nil {
		t.Fatal(err)
	}
}

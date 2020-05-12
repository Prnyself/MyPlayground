package dir

import "testing"

func TestStat(t *testing.T) {
	if err := Stat("../dir"); err != nil {
		t.Fatal(err)
	}
	if err := Stat("."); err != nil {
		t.Fatal(err)
	}
}

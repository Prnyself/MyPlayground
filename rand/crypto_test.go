package rand

import "testing"

func TestCrypto(t *testing.T) {
	if err := Crypto(); err != nil {
		t.Fatal(err)
	}
}

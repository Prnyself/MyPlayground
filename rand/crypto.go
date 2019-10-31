package rand

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Crypto() error {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return err
	}
	fmt.Printf("%x, len: %d\n", buf, len(buf))
	fmt.Printf("%s\n", hex.EncodeToString(buf))
	return nil
}

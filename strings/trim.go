package strings

import (
	"fmt"
	"strings"
)

func TrimPrefix(s, p string) error {
	res := strings.TrimPrefix(s, p)
	fmt.Println(res)
	return nil
}

func TrimLeft(s, l string) error {
	res := strings.TrimLeft(s, l)
	fmt.Println(res)
	return nil
}

package strings

import (
	"fmt"
	"strings"
)

func Split(str, sep string) error {
	fmt.Println(strings.Split(str, sep))
	return nil
}

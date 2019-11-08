package os

import (
	"fmt"
	"os"
)

func Stat(input string) error {
	info, err := os.Stat(input)
	if err != nil {
		return err
	}
	fmt.Println(info.ModTime())
	fmt.Println(info.ModTime().Unix())
	return nil
}

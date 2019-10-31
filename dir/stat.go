package dir

import (
	"fmt"
	"os"
)

func Stat(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	fmt.Printf("now path: %s, name: %s, size: %v, dir? %v\n",
		path, info.Name(), info.Size(), info.IsDir())
	return nil
}

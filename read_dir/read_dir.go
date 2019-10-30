package read_dir

import (
	"fmt"
	"io/ioutil"
)

func Show(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Printf("name: %s, size: %v, dir? %v\n",
			file.Name(), file.Size(), file.IsDir())
	}
	return nil
}

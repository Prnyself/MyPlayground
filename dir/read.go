package dir

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Read(path string) error {
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

func ReadRecursively(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("now path: %s, name: %s, size: %v, dir? %v\n",
			path, info.Name(), info.Size(), info.IsDir())
		return nil
	})
	return err
}

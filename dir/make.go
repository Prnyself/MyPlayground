package dir

import (
	"os"
	"path/filepath"
	"strings"
)

func Make(path string) error {
	if !strings.HasSuffix(path, string(os.PathSeparator)) {
		path = filepath.Dir(path)
	}
	return os.MkdirAll(path, 0755)
}

func CreateFile(path string) error {
	_, err := os.Create(path)
	return err
}

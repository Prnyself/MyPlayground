package dir

import "path/filepath"

func Glob(path string) ([]string, error) {
	return filepath.Glob(path)
}

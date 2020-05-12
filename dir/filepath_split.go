package dir

import (
	"path/filepath"
	"strings"
)

func Split(input string) (wd, path string, err error) {
	var absPath string
	absPath, err = filepath.Abs(input)
	if err != nil {
		return
	}
	if strings.HasSuffix(input, string(filepath.Separator)) || strings.HasSuffix(input, ".") {
		absPath = absPath + "/"
	}
	wd, path = filepath.Split(absPath)
	return
}

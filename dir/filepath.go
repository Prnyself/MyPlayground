package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

func Path(input string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Printf("work dir: <%v>\n", pwd)
	fmt.Printf("input: <%v>\n", input)

	dir := filepath.Dir(input)
	fmt.Printf("filepath.Dir: <%v>\n", dir)

	absDir, err := filepath.Abs(input)
	if err != nil {
		return err
	}
	fmt.Printf("filepath.Abs: <%v>\n", absDir)

	base := filepath.Base(input)
	fmt.Printf("filepath.Base: <%v>\n", base)

	fmt.Printf("filepath.Base abs: <%v>\n", filepath.Base(absDir))

	relPath, err := filepath.Rel(dir, input)
	if err != nil {
		return err
	}
	fmt.Printf("filepath.Rel: base <%s>, target <%s>; res: <%s>\n", dir, input, relPath)

	spDir, spFile := filepath.Split(input)
	fmt.Printf("filepath.Split: input <%s>, dir <%s>, file <%s>\n", input, spDir, spFile)
	fmt.Println()
	return nil
}

func AbsAndDir(path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	dir := filepath.Dir(abs)
	fmt.Printf("input path: <%s>, abs: <%s>, dir: <%s>\n", path, abs, dir)
	return nil
}

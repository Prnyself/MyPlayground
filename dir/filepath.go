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
	fmt.Println()
	return nil
}

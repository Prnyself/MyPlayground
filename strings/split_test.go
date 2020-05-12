package strings

import (
	"fmt"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	if err := Split("127.0.0.1", ":"); err != nil {
		t.Fatal(err)
	}

	s := strings.SplitAfter("/a///b", "/")
	fmt.Println("len:", len(s), "res:", s, "dir:", s[0:len(s)-1], "file:", s[len(s)-1])
}

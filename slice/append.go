package slice

import "fmt"

func KeepAppend() error {
	ori := make([]int, 3)
	for i := 1; i < 100; i++ {
		ori[0] = 0
		ori = ori[1:]
		ori = append(ori, i)
		fmt.Println(ori)
	}
	return nil
}

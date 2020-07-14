package slice

import "testing"

func TestKeepAppend(t *testing.T) {
	KeepAppend()
}

func BenchmarkAppendMake0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]string, 0)
		for i := 0; i < 1000; i++ {
			arr = append(arr, "a")
		}
	}
}

func BenchmarkAppendMake1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]string, 0, 1000)
		for i := 0; i < 1000; i++ {
			arr = append(arr, "a")
		}
	}
}

func BenchmarkAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := make([]string, 1000)
		for i := 0; i < 1000; i++ {
			arr[i] = "a"
		}
	}
}

func BenchmarkAppendVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var arr []string
		for i := 0; i < 1000; i++ {
			arr = append(arr, "a")
		}
	}
}

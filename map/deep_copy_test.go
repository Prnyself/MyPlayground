package _map

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/uuid"
)

var n = 20

func BenchmarkDeepCopy(b *testing.B) {
	m := conductMap(n)
	for i := 0; i < b.N; i++ {
		DeepCopy(m)
	}
}

func BenchmarkManualCopy(b *testing.B) {
	m := conductMap(n)
	for i := 0; i < b.N; i++ {
		ManualCopy(m)
	}
}

func BenchmarkMarshallCopy(b *testing.B) {
	m := conductMap(n)
	for i := 0; i < b.N; i++ {
		MarshallCopy(m)
	}
}

func conductMap(size int) map[string]State {
	rand.Seed(time.Now().Unix())
	res := make(map[string]State, size)
	for i := 0; i < size; i++ {
		id := uuid.New().String()
		num := rand.Int63()
		res[id] = State{
			Name:   id,
			Status: id,
			Done:   num,
			Total:  num,
		}
	}
	return res
}

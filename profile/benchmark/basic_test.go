package benchmark

import (
	"sort"
	"testing"
)

//go test -run xyz -bench .

//go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out
//go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out -gcflags -m=2

func BenchmarkBubbleSort(b *testing.B) {
	els := getElements(100)
	for i := 0; i < b.N; i++ {
		bubbleSort(els)
	}
}

func BenchmarkSort(b *testing.B) {
	els := getElements(100)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

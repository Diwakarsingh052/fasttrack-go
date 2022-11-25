package main

import "testing"

func BenchmarkAlgorithmOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algOne()
	}
}

//go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out
//go tool pprof p.out
//weblist algOne

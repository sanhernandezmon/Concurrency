package main

import (
	"testing"
)

const samples = 10000000

func BenchmarkPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PI(samples)
	}
}

func BenchmarkConcurrentPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentPi(samples, 4)
	}
}

func BenchmarkParallelPI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParallelPI(samples)
	}
}

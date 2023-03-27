package main

import (
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(60)
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(40)
	}
}

func BenchmarkThread(t *testing.B) {
	t.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fibonacci(40)
		}
	})

}

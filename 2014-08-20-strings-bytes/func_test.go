package main

import (
	"testing"
)

func BenchmarkCopyString(b *testing.B) {
	from := "foobar"
	var to string
	for i := 0; i < b.N; i++ {
		to = from
	}
	_ = to
}

func BenchmarkCopyBytes(b *testing.B) {
	from := []byte("foobar")
	var to []byte
	for i := 0; i < b.N; i++ {
		to = from
	}
	_ = to
}

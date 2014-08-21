package main

import (
	"bufio"
	"io"
	"testing"
)

type DummyReader int

func (r DummyReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		if r%10 == 9 {
			p[i] = '\n'
		} else {
			p[i] = 'a'+byte(r)
		}
		r++
	}
	return len(p), nil
}

var r DummyReader

var m = map[string]int{"abcdefghi": 1}

func BenchmarkMapString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		n := 0
		scanner := bufio.NewScanner(io.LimitReader(r, 1<<10))
		for scanner.Scan() {
			str := scanner.Text()
			n += m[str]
		}
	}
}

func BenchmarkMapBytes(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		n := 0
		scanner := bufio.NewScanner(io.LimitReader(r, 1<<10))
		for scanner.Scan() {
			data := scanner.Bytes()
			n += m[string(data)]
		}
	}
}

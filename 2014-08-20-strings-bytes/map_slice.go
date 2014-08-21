// +build none

package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(strings.NewReader("foo\nbar\nbaz\n"))
	// START OMIT
	m := map[[]byte]int{
		[]byte("foo"): 1,
		[]byte("baz"): 3,
	}
	for scanner.Scan() {
		data := scanner.Bytes() // no more copy!
		val := m[data]
		fmt.Printf("%q: %d\n", data, val)
	}
	// END OMIT
}

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
	m := map[string]int{
		"foo": 1,
		"baz": 3,
	}
	for scanner.Scan() {
		str := scanner.Text() // copy :(
		val := m[str]
		fmt.Printf("%q: %d\n", str, val)
	}
	// END OMIT
}

// +build none

package main

import (
	"io"
	"os"
)

func main() {
	WriteFoo(os.Stdout)
}

//START OMIT
func WriteFoo(w io.Writer) {
	str := "foo"
	w.Write([]byte(str))
}
//END OMIT

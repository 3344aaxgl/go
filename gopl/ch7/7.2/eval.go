package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	var w io.Writer
	w = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Printf("%T\n", rw)
	f := w.(*os.File)
	fmt.Printf("%T\n", f)
}

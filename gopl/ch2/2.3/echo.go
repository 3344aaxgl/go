package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")
var slice = []string{"1", "2", "3", "4", "5"}

func test(arr [5]int) {
	for i := range arr[0:] {
		arr[i] += i
	}
}

func main() {
	flag.Parse()
	fmt.Print(flag.Args())
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

	arr := [5]int{0, 1, 2, 3, 4}
	test(arr)
	fmt.Print(arr)
}

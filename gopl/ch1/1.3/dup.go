package main

import(
	"fmt"
	"bufio"
	"os"
)

func dup1(){
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan(){
		counts[input.Text()]++;
	}

	for line, n := range counts{
		fmt.Printf("line %s: count:%d\n", line, n)
	}
}

func dup2(){
	counts := make(map[string]int)
	for _, arg := range os.Args[1:]{
		if len(arg) > 0 {
			f, error := os.Open(os.Args[1])
			if error != nil {
				countlines(f, counts)
			}
		} else{
			countlines(os.Stdin, counts)
		}
	}
	
}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main(){


}
package main

import(
	"fmt"
	"os"
)

func echo1(){
	var s, step string
	for i := 0; i < len(os.Args); i++{
		s += step + os.Args[i]
		step = " "
	}
	fmt.Println(s)
}

func echo2(){
	var s, step string
	for _, arg := range os.Args{
		s += step + arg
		step = " "
	}
	fmt.Println(s)
}

func main(){
	echo1()
	fmt.Println()
	echo2()
}
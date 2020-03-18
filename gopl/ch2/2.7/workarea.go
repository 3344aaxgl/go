package main

import "fmt"

func f(){}
var g= "g"

func j() int{return 1}
func k(x int) int{return x}

func main(){
	f := "f"
	fmt.Println(f)
	fmt.Println(g)
	//fmt.Println(h)
	fmt.Println("-----------------")

	x := "hello!"
	for i := 0; i< len(x); i++{
		x := x[i]
		if x != '!'{
			x := x + 'A' -'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println("\n-----------------")
	for _, x:= range x{
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
	fmt.Println("\n-----------------")
	if x:= j(); x == 0{
		fmt.Println(x)
	} else if y:= k(x); x == y{
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
}
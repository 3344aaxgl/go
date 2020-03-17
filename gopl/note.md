# 3 在将数组作为函数参数时，函数接收的时数组的副本，并没有和C/C++一样，将数组退化成指针。

```
package main

import "fmt"

func test(arr [5]int) {
	for i := range arr[0:] {
		arr[i] += i
	}
}

func main() {
	arr := [5]int{0, 1, 2, 3, 4}
	test(arr)
	fmt.Print(arr) //[0 1 2 3 4]
}

```
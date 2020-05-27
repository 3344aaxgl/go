# 1. 每次调用new函数都是返回一个新的变量的地址,除非类型为空。但并不一定是在堆上创建（PS:只是一个语法糖，因为拥有GC机制，所以不需要开发人员决定变量创建的位置）

# 2. 对于在包一级声明的变量来说，他们的生命周期和整个程序的运行周期是一致的。局部变量的生命周期则是动态的：每次从创建一个新变量的声明语句开始，直到该变量不在被引用为止，然后变量的存储空间可能被回收。(PS:包一级的变量首字母大写相当于全局变量，可以在其他包中引用，而首字母小写则相当于static变量，只能在包内引用。局部变量是创建在堆还是栈上取决于是否逃逸)

```go
var global *int

//x必须在堆上创建，因为函数结束后，仍可以通过gobal访问到，称之为“逃逸”
func f() {
    var x int
    x = 1
    global = &x
}

//*y没有“逃逸”，由编译器选择创建的位置
func g() {
    y := new(int)
    *y = 1
}
```

PS:这里与C/C++有很大的不同，在《Effective C++》中，建议我们不要在函数中返回函数内声明的局部变量的指针或者引用，这是因为局部变量是分配在栈上，函数结束后就消亡了,可以直接返回局部变量对象，这里涉及到编译器的**RVO**(Return Value Optimization)。本质上还是因为golang拥有GC机制，可以由编译器决定变量的创建位置，当然也会有一些性能的损失。

# 3. 在将数组作为函数参数时，函数接收的时数组的副本，并没有和C/C++一样，将数组退化成指针。

```go
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

# 4. slice

slice的结构类似于c++中的string，一个指向数组的指针，一个存储数据的长度，一个已分配空间的长度，也有自动扩容的功能。拷贝操作是深拷贝。

```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```

# 5. Map

Map是底层为哈希表的键值对容器，其中key必须是支持==比较运算符的数据类型.迭代顺序是不确定的，并且每一次遍历的顺序都不同

# 6. 结构体

如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的。拥有匿名成员可以直接访问匿名成员的属性而不用给出完整路径

# 7. 函数

Go和C一样，不支持重载。如果两个函数参数列表和返回值列表中的变量类型一一对应，那么这两个函数被认为有相同的类型和标识符。go的函数可以有多个返回值。

# 8. 函数值

函数值类似函数指针，可以作为变量来创建，赋值，可以与nil比较，但函数值之间不可以比较，也不能当作map的key

# 9. 匿名函数

匿名函数类似lambda表达式.下面这个例子很有意思。首先类似lambda表达式，可以访问外部变量，但lambda表达式需要指明可以访问的外部变量，是传值，还是引用。函数值记录的变量的地址。然后是这个x变量没有在第一次调用函数值后被释放，变量的生命周期不由它的作用域决定。这也说明函数值不仅仅是一串代码，还记录了状态，也是函数值属于引用类型且不可比较

```go
package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) //1
	fmt.Println(f()) //4
	fmt.Println(f()) //9
	fmt.Println(f()) //16
}
```

函数值记录的是外部变量的地址导致的一个坑

```go
	var rmdirs []func()
	for _, d := range tempDirs() {
		dir := d               // NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	// ...do some work…
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
```

这里必须在for循环内对变量d进行一次拷贝，os.RemoveAll中记录的是dir的地址，而其中存储的是最后一个目录

# 10. 可变参数

类似C/C++可变参数

# 11. Deferred函数

在调用普通函数或方法前加上关键字defer，跟在defer后面的函数就会被延迟执行。多条defer语句时，执行顺序与声明顺序相反。作用有点像析构函数。《UNIX环境高级编程》中也提到过aexit函数，可以注册程序结束时调用的函数。参数记录的也是地址

# 12. Panic异常

当panic异常发生时，程序会中断运行，并立即执行在该goroutine中被延迟的函数（defer机制）

# 13. Recover捕获异常

如果在deferred函数中调用了内置函数recover，并且定义该defer语句的函数发生了panic异
常，recover会使程序从panic中恢复，并返回panic value。导致panic异常的函数不会继续运
行，但能正常返回。在未发生panic时调用recover，recover会返回nil。

# 14. 方法

在函数声明时，在其名字之前放上一个变量（接收器），即是一个方法。如果接收器本身比较大或者需要更新接收器本身时，就需要用其指针而不是对象来声明方法.包含结构体的同时，也得到了该结构体的方法，匿名类型的方法可以直接使用。作用有点像继承

# 15. 封装

Go语言只有一种控制可见性的手段：大写首字母的标识符会从定义它们的包中被导出，小写
字母的则不会。这种限制包内成员的方式同样适用于struct或者一个类型的方法。

# 16. 接口


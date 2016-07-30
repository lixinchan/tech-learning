// basic_func
package main

import (
	"fmt"
)

func MyPrint(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "it's int")
		case string:
			fmt.Println(arg, "it's string")
		case float64:
			fmt.Println(arg, "it's float64")
		default:
			fmt.Println(arg, "it's unknown type")
		}
	}
}

func main() {

	// 匿名函数&闭包
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Println("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()

	// 不定参数
	var v1 int = 3
	var v2 string = "hello"
	var v3 float64 = 4.5
	var v4 complex64 = complex(3, 3.5)
	MyPrint(v1, v2, v3, v4)

}

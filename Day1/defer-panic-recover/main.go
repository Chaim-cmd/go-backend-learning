package main

import "fmt"

//defer：LIFO顺序（后进先出)
func deferDemo() {
	fmt.Println("开始")
	defer fmt.Println("defer3")
	defer fmt.Println("defer2")
	defer fmt.Println("defer1")
	fmt.Println("函数执行中")
}

//recover 必须在defer 里调用才能捕获 panic
func safeDiv(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("捕获到panic: %v", r)
		}
	}()
	result = a / b
	return
}

func main() {
	deferDemo()

	result, err := safeDiv(10, 2)
	fmt.Println(result, err)
	result1, err1 := safeDiv(10, 0)
	fmt.Println(result1, err1)
}

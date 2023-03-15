package main

import "fmt"

func main() {
	//one()
	two_1()
}

func one() {
	const pi = 3.14159
	const (
		e   = 2.7182818
		pi1 = 3.1415926
	)
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)
}

func one_1() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
}

func two_1() {
	var cat int = 1
	var str string = "ms的go教程"
	fmt.Printf("%p %p", &cat, &str)
	fmt.Println()

	var room int = 10 // room房间 里面放的 变量10
	var ptr = &room   // 门牌号px  指针  0xc00000a0a8

	fmt.Printf("%p\n", &room) // 变量的内存地址 0xc00000a0a8
	fmt.Println()

	fmt.Printf("%T, %p\n", ptr, ptr) // *int, 0xc00000a0a8
	fmt.Println()

	fmt.Println("指针地址", ptr)          // 0xc00000a0a8
	fmt.Println("指针地址代表的值", *ptr) // 10
}

func two_2() {

}

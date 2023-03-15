package main

import (
	"fmt"
	"net"
)

func main() {

	//one()
	//two()
	//three()
	//four()
	//five_one()
	//five_two()
	five_three()
}

func one() {
	a := 100
	b := 200
	var c int
	c = b
	b = a
	a = c
	fmt.Printf("a=%d,b=%d", a, b)
}

func two() {
	a := 100
	b := 100

	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Printf("a=%d,b=%d", a, b)
}

func three() {
	a := 100
	b := 200
	a, b = b, a
	fmt.Printf("a=%d,b=%d", a, b)
}

func four() {
	dial, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println(dial)

	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	conn1, _ := net.Dial("tcp", "127.0.0.1:8080")
	fmt.Println(conn)
	fmt.Println(conn1)
}

// ## 5.1 局部变量
func five_one() {
	var a = 3
	var b = 4
	c := a + b
	fmt.Printf("a =%d,b=%d.c=%d\n", a, b, c)
}

// ## 5.2 全局变量
var c int
var bb float32 = 3.14

func five_two() {
	var a, b int
	a = 3
	b = 4
	c = a + b
	fmt.Printf("a =%d,b=%d,c=%d\n", a, b, c)

	bb := 3
	fmt.Println(bb)
}

var a = 13

func five_three() {
	var a = 3
	var b = 4
	fmt.Printf("main() 函数中 a = %d\n", a)
	fmt.Printf("main() 函数中 b = %d\n", b)
	c := sum(a, b)
	fmt.Printf("main() 函数中 c = %d\n", c)
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

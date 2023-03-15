package main

import "fmt"

var age int
var level = 1
var (
	a int
	b string
	c []float32
)

func main() {
	one()

}

func one() {
	fmt.Println(age)
	fmt.Printf("%T", level)
	fmt.Println()
	fmt.Printf("%d,%s,%f", a, b, c)
	fmt.Println()
	aa := 1
	fmt.Println(aa)
}

func two() {

}

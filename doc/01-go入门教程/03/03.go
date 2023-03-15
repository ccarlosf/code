package main

import "fmt"

func main() {
	//two()
}

func two() {
	floatStr1 := 3.2
	fmt.Printf("%.2f\n", floatStr1)

	var f1 float32 = 1 << 24
	fmt.Println(f1 == f1+1)

	var e = .71828
	var f2 = 1.
	fmt.Printf("%.5f,%.1f", e, f2)
	fmt.Println()

	var avogadro = 6.02214129e23 // 阿伏伽德罗常数
	var planck = 6.62606957e-34  // 普朗克常数
	fmt.Printf("%f,%.35f", avogadro, planck)
}

func three() {

}

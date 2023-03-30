package main

import "fmt"

func main() {
	//one()
	//two()
	//three()
	//three_one()
	//three_2()
	//three_3()
	three_4()
	//five()
}

var arr [3]int

func one() {
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
}

func two() {
	/*for index, value := range arr {
		fmt.Printf("索引:%d,值：%d \n", index, value)
	}*/

	//var arr [3]int = [3]int{1, 2, 3}
	//var arr [3]int = [3]int{1, 3}
	//arr:=[3]int{1,2,3}
	//arr:=[...]int{1,2,3}

	//var arr [3]int
	//arr[0] = 5
	//arr[1] = 6
	//arr[2] = 7

	type arr3 [3]int
	var arr arr3
	arr[0] = 2
	for index, value := range arr {
		fmt.Printf("索引:%d,值：%d \n", index, value)
	}

}

func three() {
	var a = []int{1, 2, 3}
	fmt.Println(a, a[1:2])

	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	fmt.Println(highRiseBuilding[10:15])
	fmt.Println(highRiseBuilding[20:])
	fmt.Println(highRiseBuilding[:2])

}

func three_one() {
	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)

	// 追加一个元素
	strList = append(strList, "ms的go教程")
	fmt.Println(strList)
}

func three_2() {
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))
}

func three_3() {
	var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	myslice := numbers4[4:6]
	fmt.Printf("myslice为 %d, 其长度为: %d\n", myslice, len(myslice))
	myslice = myslice[:cap(myslice)]
	fmt.Printf("myslice的第四个元素为: %d", myslice[3])
}

func three_4() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1)
	copy(slice1, slice2)

	const elementCount = 1000
	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	refData := srcData
	copData := make([]int, elementCount)
	copy(copData, srcData)
	srcData[0] = 999
	fmt.Println(refData[0])
	fmt.Println(copData[0], copData[elementCount-1])
	copy(copData, srcData[4:6])
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copData[i])
	}
}

func five() {
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapAssigned = mapLit
	mapAssigned["two"] = 3
	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}

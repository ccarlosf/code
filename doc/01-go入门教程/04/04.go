package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//one()
	//two_1()
	//two_2()
	//two_3()
	//two_4()
	//five()
	//eight_1()
	//eight_2()
	nine()
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
	var cat = 1
	var str = "ms的go教程"
	fmt.Printf("%p %p", &cat, &str)
	fmt.Println()

	var room = 10   // room房间 里面放的 变量10
	var ptr = &room // 门牌号px  指针  0xc00000a0a8

	fmt.Printf("%p\n", &room) // 变量的内存地址 0xc00000a0a8
	fmt.Println()

	fmt.Printf("%T, %p\n", ptr, ptr) // *int, 0xc00000a0a8
	fmt.Println()

	fmt.Println("指针地址", ptr)      // 0xc00000a0a8
	fmt.Println("指针地址代表的值", *ptr) // 10
}

func two_2() {

	var num = 10
	modifyFromPoint(num)
	fmt.Println("未使用指针，方法外", num)

	var num2 = 22
	newModifyFromPoint(&num2) // 传入指针
	fmt.Println("使用指针 方法外", num2)

}

func modifyFromPoint(num int) {
	num = 10000
	fmt.Println("未使用指针，方法内:", num)

}

func newModifyFromPoint(ptr *int) {
	*ptr = 2000
	fmt.Println("未使用指针，方法内:", *ptr)
}

func two_3() {
	str := new(string)
	*str = "ms的go教程Go语言教程"
	fmt.Println(*str)
}

func two_4() {
	var mode = flag.String("mode", "", "fast模式能让程序运行的更快")

	flag.Parse()
	fmt.Println(*mode)
}

type NewInt int
type IntAlias = int

func five() {
	var a NewInt
	fmt.Printf("a type: %T\n", a)

	var a2 IntAlias
	fmt.Printf("a2 type: %T\n", a2)

}

func eight_1() {
	newStr1 := "1"
	intValue, _ := strconv.Atoi(newStr1)
	fmt.Printf("%T,%d\n", intValue, intValue)

	intValue2 := 1
	sttValue := strconv.Itoa(intValue2)
	fmt.Printf("%T,%s\n", sttValue, sttValue)
}

func eight_2() {
	string3 := "3.1415926"
	f, _ := strconv.ParseFloat(string3, 32)
	fmt.Printf("%T,%f\n", f, f)

	floatValue := 3.1415926
	formatFloat := strconv.FormatFloat(floatValue, 'f', 2, 64)
	fmt.Printf("%T,%s", formatFloat, formatFloat)
}

var level = 1
var ex = 0

func nine() {
	fmt.Println("请输入你的角色名字")
	//捕获标准输入，并转换为字符串
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//删除最后的\n
	name := input[:len(input)-1]
	fmt.Printf("角色创建成功,%s,欢迎你来到张三游戏,目前角色等级%d \n", name, level)
	s := `你遇到了一个怪物，请选择是战斗还是逃跑?
   	1.战斗
   	2.逃跑`
	fmt.Printf("%s \n", s)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		selector := input[:len(input)-1]
		switch selector {
		case "1":
			ex += 10
			fmt.Printf("杀死了怪物，获得了%d经验 \n", ex)
			computeLevel()
			fmt.Printf("您现在的等级为%d \n", level)
		case "2":
			fmt.Printf("你选择了逃跑\n")
			fmt.Printf("%s \n", s)
		case "exit":
			fmt.Println("你退出了游戏")
			//退出
			os.Exit(1)
		default:
			fmt.Println("你的输入我不认识，请重新输入")
		}
	}
}

func computeLevel() {
	if ex < 20 {
		level = 1
	} else if ex < 40 {
		level = 2
	} else if ex < 200 {
		level = 3
	} else {
		level = 4
	}
}

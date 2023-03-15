package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func main() {
	//two()
	//three()
	//four()
	//five()
	//five_one()
	//five_1()
	//five_2()
	//five_3()
	//five_4()
	//six_1()
	seven()
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
	var a = 10
	if a > 11 && a < 30 {
		fmt.Println("正确")
	} else {
		fmt.Println("错误")
	}

	if a > 5 || a < 30 {
		fmt.Println("正确")
	} else {
		fmt.Println("错误")
	}
}

func four() {
	//使用单引号 表示一个字符
	//var ch byte = 'A'
	//在 ASCII 码表中，A 的值是 65,也可以这么定义
	//var ch byte = 65
	//65使用十六进制表示是41，所以也可以这么定义 \x 总是紧跟着长度为 2 的 16 进制数
	var ch byte = '\x41'
	//65的八进制表示是101，所以使用八进制定义 \后面紧跟着长度为 3 的八进制数
	//var ch byte = '\101'

	fmt.Printf("%c", ch)
}

func five() {
	var str = "ms的go教程\nGo大法好"
	fmt.Println(str)

	fmt.Println(`\t ms的go教程Go大法好`)  // \t ms的go教程Go大法好
	fmt.Println(`\t ms的go教程
 Go大法好`) //使用反引号 可以进行字符串换行
	//反引号一般用在 需要将内容进行原样输出的时候 使用

	var mystr01 = "hello"
	var mystr02 = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("myStr01: %s\n", mystr01)
	fmt.Printf("myStr02: %s", mystr02)

	//中文三字节，字母一个字节
	var myStr01 = "hello,ms的go教程"
	fmt.Printf("mystr01: %d\n", len(myStr01))
}

func five_one() {
	str3 := "hello"
	str4 := "你好"
	fmt.Println(len(str3))
	fmt.Println(len(str4))
	fmt.Println(utf8.RuneCountInString(str4))

	str1 := "你好,"
	str2 := "ms的go教程"

	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)
	fmt.Println(stringBuilder.String())
}

func five_1() {
	var myStr01 = "hello,ms的go教程"
	fmt.Println(string([]rune(myStr01)[06]))
}

func five_2() {
	var str1 = "hello"
	var str2 = "hello,ms的go教程"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("ascii: %c %d\n", str1[i], str1[i])
	}
	for _, s := range str1 {
		fmt.Printf("unicode: %c %d\n ", s, s)
	}
	// 中文只能用 for range
	for _, s := range str2 {
		fmt.Printf("unicode: %c %d\n ", s, s)
	}
}

func five_3() {
	str1 := "你好,"
	str2 := "ms的go教程"
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(str1)
	stringBuilder.WriteString(str2)
	// Sprint 以字符串形式返回
	result := fmt.Sprintf(stringBuilder.String())
	fmt.Println(result)
}

func five_4() {
	// 查找
	tracer := "张三来了,张三bye bye"

	comma := strings.Index(tracer, ",")
	fmt.Println(",所在的位置:", comma)
	fmt.Println(tracer[comma+1:]) // 张三bye bye

	add := strings.Index(tracer, "+")
	fmt.Println("+所在的位置:", add) // +所在的位置: -1

	pos := strings.Index(tracer[comma:], "张三")
	fmt.Println("张三，所在的位置", pos) // 张三，所在的位置 1

	fmt.Println(comma, pos, tracer[5+pos:]) // 12 1 张三bye bye
}

func six() {
	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", a, a)
	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)
	// 将常量保存为float32类型
	var c float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(c))
}

func six_1() {
	s1 := "localhost:8080"
	fmt.Println(s1)

	strByte := []byte(s1)

	strByte[len(s1)-1] = '1'
	fmt.Println(strByte)

	s2 := string(strByte)
	fmt.Println(s2)
}

func seven() {
	str1 := "Hello, ms的go教程Java教程"
	source := "Java"
	target := "Go"

	index := strings.Index(str1, source)
	sourceLength := len(source)

	start := str1[:index]
	end := str1[index+sourceLength:]

	var stringBuilder bytes.Buffer
	stringBuilder.WriteString(start)
	stringBuilder.WriteString(target)
	stringBuilder.WriteString(end)
	fmt.Println(stringBuilder.String())

}

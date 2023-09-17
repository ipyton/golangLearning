package main

import (
	"GoLearning/c3/utils"
	"fmt"
	"strconv"
	"unsafe"
)

func different_type() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myCar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("%T , %q", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("%T , %q", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("%T , %q", str, str)

	str = fmt.Sprintf("%c", myCar)
	fmt.Printf("%T , %q", str, str)

	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("--------%T %q\n------", num3, num3)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("%T %q", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("%T %q", str, str)

	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("%T %q", str, str)

}

func size() {
	var num1 = 99.99
	num2 := 88
	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
}

func char() {
	var c1 byte = 'a'
	c2 := 'b'
	fmt.Printf("%c %c", c1, c2)
	c3 := 22269
	fmt.Printf("%c", c3)

}
func module() {
	fmt.Printf(utils.HeroName)
}

func floatDemo() {

}

func main() {
	different_type()
	size()
	char()
	module()

}

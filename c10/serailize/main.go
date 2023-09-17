package main

import (
	json2 "encoding/json"
	"fmt"
)

// convert
type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name""`
	Age   int    `json:"age""`
	Skill string `json:"skill"`
}

type Student struct {
}
type AA struct {
	Name string
	age  int
}

func (a *AA) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *AA) hello() {
	fmt.Println("A hello", a.Name)
}

type BB struct {
	AA
	Name string
}

func (b *BB) SayOk() {
	fmt.Println("B SayOk", b.Name)
}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是 bool 类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是 float32 类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是 float64 类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是 整数 类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是 string 类型，值是%v\n", index, x)
		case Student:
			fmt.Printf("第%v个参数是 Student 类型，值是%v\n", index, x)
		case *Student:
			fmt.Printf("第%v个参数是 *Student 类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数是  类型 不确定，值是%v\n", index, x)

		}
	}
}

func main() {
	var a A
	var b B
	a = A(b)
	fmt.Println(a)

	mon := Monster{"cow", 500, "fan"}
	json, _ := json2.Marshal(mon)
	fmt.Println(string(json))

	var v BB
	v.Name = "jack" // ok
	v.AA.Name = "scott"
	v.age = 100  //ok
	v.SayOk()    // B SayOk  jack
	v.AA.SayOk() //  A SayOk scott
	v.hello()    //  A hello ? "jack" 还是 "scott"
}

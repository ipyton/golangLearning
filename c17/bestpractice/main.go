package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"grades"`
	Sex   string
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func (s Monster) Print() {
	fmt.Println()
}

func main() {
	a := Monster{
		"jane",
		18,
		20,
		"man",
	}

	var b interface{}
	b = a

	typ := reflect.TypeOf(b)
	v := reflect.ValueOf(b)
	fmt.Println(v.Kind())
	num := v.NumField()
	for i := 0; i < num; i++ {
		fmt.Println(typ.Field(i).Tag.Get("json"))
	}
	//获取到该结构体有多少个方法
	numOfMethod := v.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//var params []reflect.Value
	//方法的排序默认是按照 函数名的排序（ASCII码）
	v.Method(1).Call(nil) //获取到第二个方法。调用它

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value //声明了 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := v.Method(0).Call(params)   //传入的参数是 []reflect.Value, 返回[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果, 返回的结果是 []reflect.Value*/

}

package main

import (
	"GoLearning/c10/modules"
	"fmt"
)

func main() {
	var stu = modules.NewStudent("tom", 99)
	fmt.Println(*stu)
	fmt.Println(stu.Name, stu.GetScore())
}

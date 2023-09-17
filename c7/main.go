package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 900
}

type Stu struct {
	Name    string
	Age     int
	Address string
}

func main() {
	map1 := make(map[int]int, 2)
	map1[1] = 90
	fmt.Println(map1)

	students := make(map[string]Stu, 10)
	stu1 := Stu{"tom", 18, "beijing"}
	stu2 := Stu{"jack", 18, "shanghai"}
	students["no1"] = stu1
	students["no2"] = stu2

	for k, v := range students {
		fmt.Println(k)
		fmt.Println(v.Name)
		fmt.Println(v.Age)
		fmt.Println(v.Address)
	}

	cities := make(map[string]string, 3)
	cities["no1"] = "beijing"
	cities["no2"] = "tianjing"

}

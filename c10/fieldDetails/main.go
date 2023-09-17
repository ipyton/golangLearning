package fieldDetails

import "fmt"

type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

type Visitor struct {
	Name string
	Age  int
}

type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
	Score [3]int
}

func (p Person) speak() {
	fmt.Println(p.Name)
}

func (p *Person) calc(n int) {

}

type Dog struct {
}

func main() {
	var p1 Person
	fmt.Println(p1)

	if p1.ptr == nil {
		fmt.Println()
	}
	p1.calc(10)

}

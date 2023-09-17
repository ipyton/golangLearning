package main

import "fmt"

func loop() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue
			}
			fmt.Println(j)
		}
	}
here:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			continue here
		}
		fmt.Println(i)
	}
}

func test(char byte) byte {
	return char
}

func branch() {
	var age int
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("lager than 18")
	} else {
	}

	var key byte
	fmt.Scanf("%c", &key)
	fmt.Println(key)
	switch test(key) + 1 {
	case 'a':
		fmt.Println("Monday")
	case 'b':
		fmt.Println("tuesday")
	case 'c':
		fmt.Println("wednesday")
	case 'd':
		fmt.Println("thursday")
	default:
		fmt.Println("error")
	}

	var n1 int32 = 51
	var n2 int32 = 20
	switch n1 {
	case n2, 10, 5:
		fmt.Println("ok1")
	case 90:
		fmt.Println("ok2~")

	}

}

func main() {
	// loop()
	branch()
}

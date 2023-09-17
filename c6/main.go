package main

import "fmt"

func str() {
	// string is immutable
	// slice of chars is mutable
	/*
			type slice struct {
			ptr *[2]int //start
			len int     // how mant elements has been saved
		    cap int     // how many elements can be saved
			}

	*/
}

func main() {
	str := "aaa"
	r := []rune(str)

	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
		fmt.Println(str[i])
	}

	var intArr [5]int = [...]int{1, 22, 33, 66, 88}
	// 1. The first method.
	var slice []int = intArr[0:3]
	//
	slice = append(slice, 11, 12, 13)
	fmt.Println(slice)
	for i, v := range slice {
		fmt.Println(i, v)
	}

	// 2. The second method.
	var slice1 = make([]int, 10)
	copy(slice1, slice)

	// 3. the difference between rune and byte(uint8, int32)

}

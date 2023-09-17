package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "d:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	// fallthrough go to next switch case clause

}

package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Println("write data", i)
	}
	close(intChan)
}

func readData(in chan int, exit chan bool) {
	for {
		v, ok := <-in
		if !ok {
			break
		}
		time.Sleep(time.Second)
		println(v)

	}
	exit <- true
	close(exit)
}

func main() {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)
	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}

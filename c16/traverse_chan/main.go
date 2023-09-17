package main

import (
	"fmt"
	"time"
)

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var myMap map[int]string
	myMap = make(map[int]string, 10)

	myMap[0] = "golang"
}

func main() {
	// after closing the channel reading from it is available
	// write only
	var channel chan<- int
	channel = make(chan int, 3)
	channel <- 20

	//read only
	var channel1 <-chan int
	num2 := <-channel1
	fmt.Println(num2)

	intChannel := make(chan int, 10)
	stringChannel := make(chan string, 5)

	for {
		select { // if can not read from a channel, we can read from the next case
		case v := <-intChannel:
			fmt.Println("read from int channel" + string(v))
			time.Sleep(time.Second)
		case v := <-stringChannel:
			fmt.Println(v)
			time.Sleep(time.Second)
		default:
			fmt.Println("can not get any one of them")
			return
		}
	}

	//Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F
	//calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller.
	//To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the
	//current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly.
	//They can also be caused by runtime errors, such as out-of-bounds array accesses.
	//Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside
	//deferred functions. During normal execution, a call to recover will return nil and have no other effect.
	//If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal
	//execution.
	//Here’s an example program that demonstrates the mechanics of panic and defer:
	test()

}

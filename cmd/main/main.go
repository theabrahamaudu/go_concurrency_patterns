package main

import (
	"fmt"
	// "time"
)

// ===== Go Routines ===== 
// func someFunc(num string) {
// 	fmt.Println(num)
// }


// func main() {

// 	go someFunc("1")
// 	go someFunc("2")
// 	go someFunc("3")

// 	time.Sleep(time.Second * 2)

// 	fmt.Println("hi")

// }


// ===== Channels =====
func main() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	msg := <- myChannel

	fmt.Println(msg)
}
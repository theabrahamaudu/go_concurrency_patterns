package main

import (
	"fmt"
	"time"
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
// func main() {
// 	myChannel := make(chan string)
// 	anotherChannel := make(chan string)

// 	go func() {
// 		myChannel <- "data"
// 	}()

// 	go func() {
// 		anotherChannel <- "cow"
// 	}()

// 	select {
// 	case msgFromMyChannel := <- myChannel:
// 		fmt.Println(msgFromMyChannel)
// 	case msgFromAnotherChannel := <- anotherChannel:
// 		fmt.Println(msgFromAnotherChannel)
// 	}
// }

// For-Select Loop Ex 1
// func main() {
// 	charChannel := make(chan string, 3)
// 	chars := []string{"a", "b", "c"}

// 	for _, s := range chars {
// 		select {
// 		case charChannel <- s:
// 		}
// 	}

// 	close(charChannel)

// 	for result := range charChannel {
// 		fmt.Println(result)
// 	}

// }

// For-Select Loop Ex 2
func main() {
	go func() {
		for {
			select {
			default:
				fmt.Println("DOING WORK")
			}
		}
	}()

	time.Sleep(time.Second * 3)
}
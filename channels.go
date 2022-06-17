// // Channels are the pipes that connect concurrent goroutines
// // You can send and recieve
// // channel <- SENDS
// // <- channel RECIEVES
// package main

// import (
// 	"fmt"
// 	"time"
// )

// //Buffered channels accept a limined number of values with out reciever needed

// // Channel synchornization is used to synch ecevution

// func worker(done chan bool) {
// 	fmt.Print("working...")
// 	time.Sleep(time.Second)
// 	fmt.Println("done")

// 	done <- true
// }
// func main() {
// 	messages := make(chan string, 2)
// 	// go func() { messages <- "ping" }()
// 	messages <- "buffered"
// 	messages <- "channel"
// 	fmt.Println(<-messages, <-messages)

// 	done := make(chan bool, 1)
// 	go worker(done)

// 	<-done
// }

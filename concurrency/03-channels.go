package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 50)
	go count("java", c)

	// a nicer way to handle messages from a channel
	/**
			for msg := range c {
				fmt.Println(msg)
			}
	    **/

	// Deadlock
	// receives the message and this is a blocking operation
	// communicating and syncing. This is an important concept to understand
	// This is because the count function is finished but the main function is waiting to receive on the channel
	// go can detect this a runtime not at compile time. It has not solved the halting problem.
	// This is why we need to close the channel
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func count(myMessage string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- myMessage
		time.Sleep(time.Millisecond * 500)
	}
	// the receiver should always close the channel not the sender
	//close(c)
}

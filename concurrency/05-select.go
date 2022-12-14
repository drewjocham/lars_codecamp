package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Every 500ms"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Every two seconds"
		}
	}()

	// remember it is blocking. What do we think is going to happen?
	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}

	/**
	  for {
	      select {
	      case msg1 := <-c1:
	          fmt.Println(msg1)
	      case msg2 := <-c2:
	          fmt.Println(msg2)
	      }
	  }
	  **/
}

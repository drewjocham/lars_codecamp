package main

import (
	"fmt"
	"time"
)

// let's run these both concurrently, what will happen and why
func main() {
	go count("java")
	count("python")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

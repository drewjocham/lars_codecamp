package main

import "fmt"

// what is going to happen? Why?
func main() {
    c := make(chan string)
    c <- "Hello"
    c <- "World"

    msg := <-c
    fmt.Println(msg)

}

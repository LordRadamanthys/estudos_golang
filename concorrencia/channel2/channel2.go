package main

import (
	"fmt"
	"time"
)

// channel é a forma de comunicao entre as go routines, é o ponto de
// sincronismo

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base
	time.Sleep(time.Second)
	c <- 3 * base
	time.Sleep(time.Second)
	c <- 4 * base

}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")
	a, b := <-c, <-c
	fmt.Println("B")
	fmt.Println(a, b)
	fmt.Println(<-c)

}

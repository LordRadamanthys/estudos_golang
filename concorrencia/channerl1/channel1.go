package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando um info para o canal
	<-ch    //recebendo uma info do canal

	ch <- 2
	fmt.Println(<-ch)
}

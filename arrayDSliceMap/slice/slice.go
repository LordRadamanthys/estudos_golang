package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3} //array
	s1 := []int{1, 2, 3}  //slice

	fmt.Println(a1, s1)

	a2 := [5]int{1, 2, 3, 4, 5}

	//slice é um treco do array

	s2 := a2[1:3]

	fmt.Println(s2)

	s3 := a2[:2]

	fmt.Println(s3)

	// vc pode imagianr um slice como tendo: um tamanho e um ponteiro para um elemento de um array
	s4 := s2[:1]
	s4[0] = 222
	fmt.Println(s4)
	fmt.Println(a2)
}

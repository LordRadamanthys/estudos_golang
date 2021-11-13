package main

import (
	"fmt"
	"strings"
)

func main() {
	s := make([]int, 10)

	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s)
	s[0] = 344
	fmt.Println(s)

	tetse := strings.ToTitle("mateus lima de matos")
	fmt.Println(tetse)
}

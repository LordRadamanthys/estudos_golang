package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("igual:", "Banana" == "Banana")
	fmt.Println("!=:", 2 != 3)
	fmt.Println("menor:", 3 < 4)
	fmt.Println("Maior:", 3 > 4)
	fmt.Println("Maior/igual:", 3 >= 4)
	fmt.Println("Maior/igual:", 3 <= 4)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"mateus"}
	p2 := Pessoa{"mateus"}

	fmt.Println("Pessoa", p1 == p2)
}

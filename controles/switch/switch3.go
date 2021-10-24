package main

import (
	"fmt"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funçao"
	default:
		return "N sei"

	}
}

func main() {
	fmt.Println(tipo("tipo"))
	fmt.Println(tipo(3.3))
	fmt.Println(tipo(8))
	fmt.Println(tipo(func() {}))
}

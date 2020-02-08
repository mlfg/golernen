package main

import (
	"fmt"
	"github.com/mlfg/golernen/aula1"
	"github.com/mlfg/golernen/aula2"
)

func main() {
	fmt.Println("hello world")

	fmt.Println(aula1.Divide(10, 2))
	fmt.Println(aula1.Divide(-30, -5))
	fmt.Println(aula1.Divide(12,60 ))
	fmt.Println(aula1.Divide(2, 4))
	fmt.Println(aula1.Divide(40, -200))
	fmt.Println(aula1.Divide(-4, 2))

	fmt.Println(aula2.Ftoc(32))
	fmt.Println(aula2.Ftoc(70))
}
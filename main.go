package main

import (
	"fmt"
	"github.com/mlfg/golernen/aula1"
	"github.com/mlfg/golernen/aula2"
	"github.com/mlfg/golernen/primes"
	"math"
	"time"
)

func main() {
	fmt.Println("this is measuringPrimeNumsTime")
	measuringPrimeNumsTime()
	fmt.Println("this is printDivResult")
	printDivResult()
	fmt.Println("this is printFahrenheit")
	printFahrenheit()
}

func measuringPrimeNumsTime(){
	T1 := time.Now()
	primes.PrintPrime1(1000)
	T2 := time.Now()

	T3 := time.Now()
	primes.PrintPrime2(1000)
	T4 := time.Now()

	elapsed1 := T2.Sub(T1)
	elapsed2 := T4.Sub(T3)

	fmt.Println(elapsed1.Milliseconds())
	fmt.Println(elapsed2.Milliseconds())
}

func printDivResult (){
	fmt.Println(aula1.Divide(10, 2))
	fmt.Println(aula1.Divide(-30, -5))
	fmt.Println(aula1.Divide(12,60 ))
	fmt.Println(aula1.Divide(2, 4))
	fmt.Println(aula1.Divide(40, -200))
	fmt.Println(aula1.Divide(-4, 2))
}

func printFahrenheit (){
	fmt.Println(aula2.Ftoc(32))
	fmt.Println(math.Round(aula2.Ftoc(73) * 100) / 100)
}
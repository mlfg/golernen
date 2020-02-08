package main

import (
	"fmt"
	"math"
	"time"
)
func main () {
	T1 := time.Now()
	printPrime1(10000)
	T2 := time.Now()

	T3 := time.Now()
	printPrime2(10000)
	T4 := time.Now()

	elapsed1 := T2.Sub(T1)
	elapsed2 := T4.Sub(T3)

	fmt.Println(elapsed1.Milliseconds())
	fmt.Println(elapsed2.Milliseconds())
}

func printPrime1 (size int) {
	primes := make(map[int]int)
	j := 0

	for i := 0; i < size; {
		if IsPrime(j) == true {
			primes[i] = j
			i++
		}
		j++
	}

	//for i := 0; i < 300; i++{
	//	fmt.Println(primes[i])
	//}
}

func printPrime2 (size int) {
	primes := make(map[int]int)
	j := 0

	for i := 0; i < size; {
		if IsPrimerSqrt(j) == true {
			primes[i] = j
			i++
		}
		j++
	}

	//for i := 0; i < 300; i++{
	//	fmt.Println(primes[i])
	//}
}

//numero divisivel por um e por ele mesmo
// +,-,*,/,%
func IsPrime (value int) bool {

	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value % i == 0 {
			return false
		}
	}

	return value > 1
}

func IsPrimerSqrt(value int) bool {

	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value >1

}





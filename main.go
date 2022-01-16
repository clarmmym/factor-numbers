package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	//fmt.Println(factor(123))
	//
	//fmt.Println(factorMultiple([]int{6, 18, 23}))

	//fmt.Print("Hello World ")
	//
	//for i := 1; i < len(os.Args); i++ {
	//
	//	a, _ := strconv.ParseInt(os.Args[i], 10, 32)
	//	fmt.Print(a, " ")
	//
	//}
	//
	//fmt.Println()

	args := os.Args[1:]
	var numbers []int

	for i := 0; i < len(args); i++ {
		number, _ := strconv.ParseInt(args[i], 10, 64)
		numbers = append(numbers, int(number))
	}

	factors := factorMultiple(numbers)

	for _, elem := range factors {
		fmt.Println(deviders(elem))
	}

}

func primeNumber(n int) []int {
	primes := make([]int, n-1)

	for i := 2; i <= n; i++ {
		primes[i-2] = i
	}

	numberSqrt := int(math.Round(math.Sqrt(float64(n))))

	element := primes[0]

	i := 1
	for element <= numberSqrt {
		for j := i; j < len(primes); j++ {
			if primes[j]%element == 0 {
				primes = append(primes[:j], primes[j+1:]...)
			}

		}

		element = primes[i]
		i++
	}

	return primes
}

type divider struct {
	prime int
	power int
}

type deviders []divider

func (d divider) String() string {
	return fmt.Sprintf("%d^%d", d.prime, d.power)
	//return strconv.Itoa(d.prime) + "^" + strconv.Itoa(d.power)
}

func (d deviders) String() string {
	s := ""

	for _, elem := range d {
		s += elem.String() + " * "
	}

	return s[:len(s)-3]
}

func factor(n int) []divider {
	var factors []divider

	if n < 0 {
		factors = append(factors, divider{-1, 1})
		n *= -1
	}

	if n == 0 || n == 1 {
		factors = append(factors, divider{n, 1})
		return factors
	}

	primes := primeNumber(n)

	for _, num := range primes {
		power := 0
		for n%num == 0 {
			power++
			n /= num
		}

		if power == 0 {
			continue
		}

		factors = append(factors, divider{num, power})
	}

	return factors
}

func factorMultiple(numbers []int) [][]divider {
	var factorMultiple [][]divider

	for _, num := range numbers {
		factorMultiple = append(factorMultiple, factor(num))
	}

	return factorMultiple
}

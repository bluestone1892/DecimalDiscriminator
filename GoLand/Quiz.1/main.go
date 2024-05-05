package main

import (
	"fmt"
	"math"
)

func isPrime(N int) bool {
	if N <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("정수 입력: ")
	fmt.Scanf("%d", &n)

	if isPrime(n) {
		fmt.Println("소수입니다.")
	} else {
		fmt.Println("소수가 아닙니다.")
	}
}

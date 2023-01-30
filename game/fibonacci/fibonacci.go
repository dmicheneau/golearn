package fibonacci

import (
	"fmt"
)

func Start() {
	var reponse int
	fmt.Printf("Enter a number:")
	fmt.Scanln(&reponse)
	if reponse < 2 {
		panic("number is too small")
	}
	fmt.Println(getFibonacci(reponse))
}

func getFibonacci(n int) []int {
	var fibo []int
	var a, b int = 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
		fibo = append(fibo, a)
	}
	return fibo
}

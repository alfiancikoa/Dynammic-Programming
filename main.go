package main

import "fmt"

// memoize the fibonacci number when we have find the solution
var memo map[int]int = map[int]int{}

func Fibonacci(n int) int {
	if memo[n] > 0 {
		return memo[n]
	} else if n == 1 || n == 2 {
		return 1
	}
	memo[n] = Fibonacci(n-1) + Fibonacci(n-2)
	return memo[n]
}

func main() {
	fmt.Println(Fibonacci(5))   //5
	fmt.Println(Fibonacci(10))  //55
	fmt.Println(Fibonacci(100)) // 3736710778780434371
}

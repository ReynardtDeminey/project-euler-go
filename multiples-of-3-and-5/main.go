package main

import "fmt"

func main() {
	res := multOf3And5(1000)
	fmt.Println(res)
}

func multOf3And5(limit int) int {
	sum := 0
	for i := 0; i < limit+1; i++ {
		if i%5 == 0 && i%3 == 0 {
			sum += i
		}
	}
	return sum
}

package main

// Problem "Oddities" https://open.kattis.com/problems/oddities

import (
	"fmt"
	"io"
)

func main() {
	var n, number int

	_, err := fmt.Scanf("%d", &n)
	if err == io.EOF {
		return
	}

	for i := 0; i < n; i++ {
		_, err := fmt.Scanf("%d", &number)
		if err == io.EOF {
			return
		}
		if (number % 2 == 0) {
			fmt.Printf("%d is even\n", number)
		} else {
			fmt.Printf("%d is odd\n", number)
		}

	}
}

package main

import (
	"fmt"
	"io"
)

func main() {
	var n int

	for {
		_, err := fmt.Scanf("%d", &n)
		if err == io.EOF {
			break
		}

		for i := 1; i <= n; i++ {
			fmt.Printf("%d Abracadabra\n", i)
		}
	}
}

package main

import "fmt"

func main() {
	var num []int
	for i := 0; i <= 10; i++ {
		num = append(num, i)
	}

	for _, n := range num {
		if n%2 == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}
}

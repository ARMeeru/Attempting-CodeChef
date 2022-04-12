// Problem Link# https://www.codechef.com/problems/FLOW001

package main

import (
	"fmt"
)

func main() {
	var t, a, b int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&a, &b)
		fmt.Println(a + b)

	}
}

// I was having trouble in this problem because I misunderstood requirement. I thought inputs will be taken first and later outputs will be shown serially. Need to think simply.

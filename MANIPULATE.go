// Problem Link# https://www.codechef.com/APRIL221D/problems/MANIPULATE

package main

import (
	"fmt"
)

func main() {
	var t, x, y int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&x, &y)

		if x == y {
			fmt.Println("YES")
		} else if x > y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

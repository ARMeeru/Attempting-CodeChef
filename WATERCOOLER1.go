// Problem Link# https://www.codechef.com/APRIL221D/problems/WATERCOOLER1

package main

import (
	"fmt"
)

func main() {
	var t, x, y, m int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&x, &y, &m)

		if x*m < y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

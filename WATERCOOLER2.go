// Problem Link# https://www.codechef.com/APRIL221D/problems/WATERCOOLER2

package main

import (
	"fmt"
)

func main() {
	var t, x, y int
	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		fmt.Scanln(&x, &y)

		if y%x != 0 {
			fmt.Println(y / x)
		} else {
			fmt.Println(y/x - 1)
		}

	}
}

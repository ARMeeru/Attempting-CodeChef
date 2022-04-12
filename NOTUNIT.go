// Problem Link# https://www.codechef.com/APRIL221D/problems/NOTUNIT

package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanln(&t)
	// Got TLE error
	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		// if even else odd
		if a%2 == 0 {
			if a+2 <= b {
				fmt.Println(a, a+2)
			} else {
				fmt.Println("-1")
			}
		} else {
			if a+3 <= b {
				if a%3 == 0 {
					fmt.Println(a, a+3)
				} else {
					fmt.Println(a+1, a+3)
				}
			} else {
				fmt.Println("-1")
			}
		}
	}

}

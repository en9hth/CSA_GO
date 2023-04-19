package main

import "fmt"

// 721递增序列
func p2() {
	var a int
	for {
		fmt.Scan(&a)
		if a == 0 {
			break
		} else {
			for i := 1; i < a+1; i++ {
				print(i, " ")
			}
			println()
		}
	}
}

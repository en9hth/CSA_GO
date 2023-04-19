package main

import "fmt"

func p4() {
	var arr [10]int
	for a := 0; a < 10; a++ {
		fmt.Scan(&arr[a])
		if arr[a] <= 0 {
			arr[a] = 1
		}
	}
	for index, val := range arr {
		fmt.Printf("X[%d] = %d\n", index, val)
	}
}

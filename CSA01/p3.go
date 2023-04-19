package main

import (
	"fmt"
	"log"
)

func p3() {
	var arr [2]float32
	if _, err := fmt.Scan(&arr[0], &arr[1]); err != nil {
		log.Fatalln(err)
	}
	if arr[0] > 0 {
		if arr[1] > 0 {
			println("Q1")
		} else if arr[1] < 0 {
			println("Q4")
		} else {
			println("Eixo X")
		}
	} else if arr[0] < 0 {
		if arr[1] > 0 {
			println("Q2")
		} else if arr[1] < 0 {
			println("Q3")
		} else {
			println("Eixo X")
		}
	} else {
		if arr[1] == 0 {
			println("Origem")
		} else {
			println("Eixo Y")
		}
	}
}

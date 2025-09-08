package main

import "fmt"

func PrintPairs(){
	for c := 0; c < 101 ; c++ {
		if c % 2 == 0 {
			fmt.Println(c)
		}
	}
}

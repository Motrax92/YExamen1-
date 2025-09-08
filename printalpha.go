package main

import "github.com/01-edu/z01"

func PrintAlpha() {
	for i := 'a'; i <= 'y'; i++ {
		for j := 'a'; j <= 'x'; j++ {
			for k := 'a'; k <= 'z'; k++ {
				if i < j && j < k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(',')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
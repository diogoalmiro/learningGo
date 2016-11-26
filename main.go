package main

import "fmt"

func main() {
	for i := 12; i >= 0; i-- {
		if i != 10 && i != 4 && i != 2 {
			fmt.Println("i =", i)
		}
	}
	recursionFTW(12)
}

func recursionFTW(i int) {
	if i != 10 && i != 4 && i != 2 {
		fmt.Println("i =", i)
	}
	if i == 0 {
		return
	}
	recursionFTW(i - 1)
}

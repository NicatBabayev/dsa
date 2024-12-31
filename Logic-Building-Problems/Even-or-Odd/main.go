package main

import "fmt"

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	if isEven(15) {
		fmt.Println("15 is even")
	} else {
		fmt.Println("15 is odd")
	}
}

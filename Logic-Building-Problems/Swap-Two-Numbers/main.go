package main

import "fmt"

func swapTwoNum(a, b int) (int, int) {
	c := a
	a = b
	b = c

	return a, b
}

func main() {
	fmt.Println(swapTwoNum(12, 222))

}

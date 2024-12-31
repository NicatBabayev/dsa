package main

import "fmt"

func sumOfNatural(n int) int {
	var res int
	for i := range n {
		i += 1
		res += i
	}
	return res
}

func main() {
	fmt.Println(sumOfNatural(5))

}

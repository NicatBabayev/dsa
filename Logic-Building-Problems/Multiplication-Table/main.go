package main

import "fmt"

func mTable(n int) {
	for i := range 10 {
		i += 1
		fmt.Printf("%d * %d = %d\n", n, i, n*i)
	}
}

func main() {
	mTable(5)
}

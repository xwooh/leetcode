package main

import (
	"fmt"
)

func main() {
	ns := []int{1, 3, 5, 2, 6, 5, 9, 4}
	quick(ns)

	fmt.Println(ns)
}

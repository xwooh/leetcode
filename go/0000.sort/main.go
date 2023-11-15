package main

import (
	"fmt"
)

func main() {
	ns := []int{5, 1, 3, 5, 2, 6, 5, 9, 4}
	// quick(ns)
	// fmt.Println(ns)

	// r := []int{}
	// topK(ns, 6, &r)
	// fmt.Println(r)

	r := merge(ns)
	fmt.Println(r)
}

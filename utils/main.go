package main

import (
	"fmt"
	"utils/sort"
)

func main() {
	xx := []int{5, 3, 5, 7, 2, 1, 85, 3, 8, 4, 63, 48, 3}
	xx = sort.QuickSort(xx)
	xx = (xx)
	fmt.Println(xx)
}

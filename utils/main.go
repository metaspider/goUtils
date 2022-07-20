package main

import (
	"fmt"
	"utils/sort"
	"utils/utils"
)

func main() {
	arr := utils.RandIntArray(-10, 10, 10)
	fmt.Println(arr)
	arr = sort.InsertSort(arr)
	fmt.Println(arr)
}

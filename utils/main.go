package main

import (
	"fmt"
	"math/rand"
	"time"
	"utils/sort"
)

func main() {
	xx := []int{}
	_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		xx = append(xx, _rand.Intn(20))
	}
	xx = sort.CountSort(xx)
	fmt.Println(xx)
}

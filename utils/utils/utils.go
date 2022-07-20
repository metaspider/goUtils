package utils

import (
	"math/rand"
	"time"
)

func RandIntArray(min, max, length int) []int {
	// 随机整数数组
	if max < min {
		min, max = max, min
	} else if max == min {
		max++
	}
	arr := make([]int, 0, length)
	offset := max - min
	_rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		arr = append(arr, _rand.Intn(offset)+min)
	}
	return arr
}

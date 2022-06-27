package sort

func BubbleSort(arr []int) []int {
	// 冒泡排序
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func QuickSort(arr []int) []int {
	// 快速排序
	if len(arr) <= 1 {
		return arr
	}
	mid := arr[0]
	left := []int{}
	right := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i] < mid {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	arr = append(QuickSort(left), mid)
	arr = append(arr, QuickSort(right)...)
	return arr
}

func SelectSort(arr []int) []int {
	// 选择排序
	for i := 0; i < len(arr); i++ {
		tmp := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[tmp] {
				tmp = j
			}
		}
		arr[i], arr[tmp] = arr[tmp], arr[i]
	}
	return arr
}

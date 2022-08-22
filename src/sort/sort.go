package sort

/*
 * (1)冒泡排序；
 * (2)选择排序；
 * (3)插入排序；
 * (4)希尔排序；
 * (5)归并排序；
 * (6)快速排序；
 * (7)基数排序；
 * (8)堆排序；
 * (9)计数排序；
 * (10)桶排序。
 */

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

func InsertSort(arr []int) []int {
	// 插入排序
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > tmp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = tmp
	}
	return arr
}

func ShellSort(arr []int) []int {
	// 希尔排序
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func merge(left, right []int) []int {
	// 归并
	arr := []int{}
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			arr = append(arr, left[0])
			left = left[1:]
		} else {
			arr = append(arr, right[0])
			right = right[1:]
		}
	}
	arr = append(arr, left...)
	arr = append(arr, right...)
	return arr
}

func MergeSort(arr []int) []int {
	// 归并排序
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}

func heapify(arr []int, i, n int) {
	tmp := arr[i]
	for k := 2*i + 1; k < n; k = 2*k + 1 {
		if k+1 < n && arr[k] < arr[k+1] {
			k++
		}
		if tmp >= arr[k] {
			break
		}
		arr[i] = arr[k]
		i = k
	}
	arr[i] = tmp
}

func HeapSort(arr []int) []int {
	// 堆排序
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
	return arr
}

func CountSort(arr []int) []int {
	// 计数排序(不支持负数)
	// 优点: 占用空间少，时间复杂度O(n+k), 空间复杂度O(n+k), 稳定, 可以用于数值范围不大的数组
	// 缺点: 占用内存大, 时间复杂度高, 只适用于小范围数据, 如0-100)
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	count := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	return res
}

func RadixSoft(arr []int) []int {
	// 基数排序(不支持负数)
	// 缺点: 占用内存大, 时间复杂度高, 只适用于小范围数据, 如0-100
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	count := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		count[arr[i]%10]++
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	res := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		res[count[arr[i]%10]-1] = arr[i]
		count[arr[i]%10]--
	}
	return res
}

func BinSort(arr []int) []int {
	// 二分查找排序
	// 优点: 时间复杂度O(nlog(n)), 空间复杂度O(1), 稳定
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for ; j >= 0; j-- {
			if arr[j] > tmp {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = tmp
	}
	return arr
}

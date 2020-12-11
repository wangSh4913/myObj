package main

import (
	"fmt"
)

/*
五种排序算法
冒泡、选择、快速、归并、插入
 */

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func selectSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return arr
}

func quickSort(start, end int, arr []int) []int {
	pivot := arr[start]
	i := start
	j := end
	for {
		if i >= j {
			break
		}
		for {
			if i < j && arr[i] < pivot {
				i++
			} else {
				break
			}
		}
		for {
			if i < j && arr[j] > pivot {
				j--
			} else {
				break
			}
		}
		if i < j && arr[i] == arr[j] {
			i++
		} else {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if i-1 > start {
		quickSort(start, i-1, arr)
	}
	if j+1 < end {
		quickSort(j+1, end, arr)
	}
	return arr
}

func merge(left []int, right []int) []int {
	index1 := 0
	index2 := 0
	tmp := make([]int, 0)
	for {
		if index1 < len(left) && index2 < len(right) {
			if left[index1] < right[index2] {
				tmp = append(tmp, left[index1])
				index1++
			} else {
				tmp = append(tmp, right[index2])
				index2++
			}
		} else {
			break
		}
	}
	tmp = append(tmp, left[index1:]...)
	tmp = append(tmp, right[index2:]...)
	return tmp
}

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := [...]int{3, 4, 1, 2, 5, 7, 6}
	fmt.Println(insertSort(arr[:]))
}

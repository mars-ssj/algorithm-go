// Package mars_sort 中实现常见的排序算法
package mars_sort

// BubbleSort 冒泡排序算法
func BubbleSort(array []int)  {
	arrayLen := len(array)
	if arrayLen <= 1 {
		return
	}
	for i := 0; i < arrayLen; i++ {
		flag := false
		for j := 0; j < arrayLen - i - 1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// InsertionSort 插入排序算法
func InsertionSort(array []int)  {
	arrayLen := len(array)
	if arrayLen <= 1 {
		return
	}
	for i := 1; i < arrayLen; i++ {
		value := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > value {
				array[j+1] = array[j]
			} else {
				break
			}
		}
		array[j+1] = value
	}
}

// SelectionSort 选择排序算法
func SelectionSort(array []int) {
	arrayLen := len(array)
	if arrayLen <= 1 {
		return
	}
	for i := 0; i < arrayLen; i++ {
		min := i
		for j := i+1; j < arrayLen; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		if i != min {
			array[i], array[min] = array[min], array[i]
		}
	}
}
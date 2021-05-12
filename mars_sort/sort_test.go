package mars_sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{2, 1, 0, -1, -3}
	BubbleSort(array)
	fmt.Println(array)
}

func TestInsertionSort(t *testing.T) {
	array := []int{2, 1, 0, -1, -3}
	InsertionSort(array)
	fmt.Println(array)
}

func TestSelectionSort(t *testing.T) {
	array := []int{2, 1, 0, -1, -3, -3, -800}
	SelectionSort(array)
	fmt.Println(array)
}
package sort

import (
	md "math/rand"
	"testing"
)

var unsort []int

func init() {
	for i := 0; i < 200; i++ {
		unsort = append(unsort, md.Int())
	}
}
func TestInsertionSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	InsertionSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("InsertionSort %v, In fact: %v", unsort, sorted)
	}
}

func TestBubbleSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	BubbleSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("InsertionSort %v, In fact: %v", unsort, sorted)
	}
}

func TestBubbleSortF(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	BubbleSortF(sorted)
	if !isSorted(sorted) {
		t.Errorf("InsertionSort %v, In fact: %v", unsort, sorted)
	}
}

func TestSelectionSort(t *testing.T) {
	sorted := make([]int, len(unsort))
	copy(sorted, unsort)
	SelectionSort(sorted)
	if !isSorted(sorted) {
		t.Errorf("InsertionSort %v, In fact: %v", unsort, sorted)
	}
}

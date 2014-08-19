package customsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	t1 := []int{2, 1, 0}
	t1e := []int{0, 1, 2}

	t2 := []int{2, 0, 1, 3, 5, 4}
	t2e := []int{0, 1, 2, 3, 4, 5}

	t3 := []int{0, 0, 3, 3, 2, 1, 7, 6}
	t3e := []int{0, 0, 1, 2, 3, 3, 6, 7}

	var sorter Sorter
	sorter = InsertionSorter{}
	t1 = sorter.Sort(t1)
	assert.Equal(t, t1, t1e)
	t2 = sorter.Sort(t2)
	assert.Equal(t, t2, t2e)
	t3 = sorter.Sort(t3)
	assert.Equal(t, t3, t3e)
}

func TestIsSorted(t *testing.T) {
	t1 := []int{2, 1, 0}
	t2 := []int{0, 1, 2}
	t3 := []int{0, 0, 1, 2}
	assert.Equal(t, IsSorted(t1), false)
	assert.Equal(t, IsSorted(t2), true)
	assert.Equal(t, IsSorted(t3), true)
}

func TestInsertionSortSwap(t *testing.T) {
	t1 := []int{2, 1, 0}
	sorter := InsertionSorter{}
	sorter.Swap(&t1, 0, 1)
	assert.Equal(t, t1, []int{1, 2, 0})
	sorter.Swap(&t1, 1, 2)
	assert.Equal(t, t1, []int{1, 0, 2})
}

func TestMergeSort(t *testing.T) {
	t1 := []int{2, 1, 0}
	t1e := []int{0, 1, 2}

	t2 := []int{2, 0, 1, 3, 5, 4}
	t2e := []int{0, 1, 2, 3, 4, 5}

	t3 := []int{0, 0, 3, 3, 2, 1, 7, 6}
	t3e := []int{0, 0, 1, 2, 3, 3, 6, 7}

	t4 := []int{1}
	t4e := []int{1}

	var sorter Sorter
	sorter = MergeSorter{}
	t1 = sorter.Sort(t1)
	assert.Equal(t, t1, t1e)
	t2 = sorter.Sort(t2)
	assert.Equal(t, t2, t2e)
	t3 = sorter.Sort(t3)
	assert.Equal(t, t3, t3e)
	t4 = sorter.Sort(t4)
	assert.Equal(t, t4, t4e)
}

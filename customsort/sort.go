package customsort

type Sorter interface {
	Sort([]int) []int
}

func IsSorted(array []int) bool {
	for i := range array {
		if i == 0 {
			continue
		}
		if array[i-1] > array[i] {
			return false
		}
	}
	return true
}

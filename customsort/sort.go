package customsort

type Sorter interface {
	Sort([]int) []int
}

func IsSorted(array []int) bool {
	var prev int = -1
	for _, e := range array {
		if prev == -1 {
			prev = e
			continue
		}
		if prev > e {
			return false
		}
	}
	return true
}

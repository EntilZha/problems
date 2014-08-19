package customsort

type InsertionSorter struct {
}

func (s InsertionSorter) Sort(arr []int) []int {
	cpy := make([]int, len(arr))
	copy(cpy, arr)
	arr = cpy
	for i := 0; i < len(arr); i++ {
		var eIndex = i
		for j := eIndex - 1; j >= 0; j-- {
			if arr[j] > arr[eIndex] {
				s.Swap(&arr, j, eIndex)
				eIndex = j
			}
		}
	}
	return arr
}

func (s InsertionSorter) Swap(arrptr *[]int, i int, j int) {
	var arr = *arrptr
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

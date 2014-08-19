package customsort

type MergeSorter struct{}

func (s MergeSorter) Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	half := len(arr) / 2
	firstHalf := s.Sort(arr[:half])
	firstCounter := 0
	secondHalf := s.Sort(arr[half:])
	secondCounter := 0
	sortedList := make([]int, 0)
	for _ = range arr {
		if len(firstHalf) == firstCounter && len(secondHalf) == secondCounter {
			return sortedList
		}

		if len(secondHalf) == secondCounter {
			sortedList = append(sortedList, firstHalf[firstCounter])
			firstCounter += 1
			continue
		}

		if len(firstHalf) == firstCounter {
			sortedList = append(sortedList, secondHalf[secondCounter])
			secondCounter += 1
			continue
		}

		if firstHalf[firstCounter] < secondHalf[secondCounter] {
			sortedList = append(sortedList, firstHalf[firstCounter])
			firstCounter += 1
		} else {
			sortedList = append(sortedList, secondHalf[secondCounter])
			secondCounter += 1
		}
	}
	return sortedList
}

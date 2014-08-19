package customsort

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var t1 = []int{2, 1, 0}
var t1e = []int{0, 1, 2}

var t2 = []int{2, 0, 1, 3, 5, 4}
var t2e = []int{0, 1, 2, 3, 4, 5}

var t3 = []int{0, 0, 3, 3, 2, 1, 7, 6}
var t3e = []int{0, 0, 1, 2, 3, 3, 6, 7}

var t4 = []int{1}
var t4e = []int{1}

func TestInsertionSort(t *testing.T) {
	var sorter Sorter
	sorter = InsertionSorter{}

	Convey("Given an unsorted list", t, func() {
		Convey("Insertion sort should sort it", func() {
			So(sorter.Sort(t1), ShouldResemble, t1e)
			So(sorter.Sort(t2), ShouldResemble, t2e)
			So(sorter.Sort(t3), ShouldResemble, t3e)
			So(sorter.Sort(t4), ShouldResemble, t4e)

		})
	})
}

func TestIsSorted(t *testing.T) {
	Convey("Given a list", t, func() {
		Convey("When it is unsorted IsSorted should return false", func() {
			So(IsSorted(t1), ShouldEqual, false)
			So(IsSorted(t3), ShouldEqual, false)
		})
		Convey("When it is sorted IsSorted should return true", func() {
			So(IsSorted(t1e), ShouldEqual, true)
			So(IsSorted(t2e), ShouldEqual, true)
		})
	})
}

func TestMergeSort(t *testing.T) {
	var sorter Sorter
	sorter = MergeSorter{}

	Convey("Given an unsorted list", t, func() {
		Convey("Merge sort should sort it", func() {
			So(sorter.Sort(t1), ShouldResemble, t1e)
			So(sorter.Sort(t2), ShouldResemble, t2e)
			So(sorter.Sort(t3), ShouldResemble, t3e)
			So(sorter.Sort(t4), ShouldResemble, t4e)
		})
	})
}

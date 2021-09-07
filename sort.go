package rhelper

import (
	"sort"
)

func SortInt32(arr []int32) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func SortInt16(arr []int16) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

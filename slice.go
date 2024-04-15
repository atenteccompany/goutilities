package goutilities

import (
	"slices"
	"sort"
)

// returns unique slice of integers from slice of integers
// this function keeps the order
func IntSet(input []int64) []int64 {
	u := make([]int64, 0, len(input))
	m := make(map[int64]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

// Intersect two slice of itegers and return intersection result.
// This is the crucial function used to achieve logical operator AND
func SliceIntersect(a []int64, b []int64) []int64 {
	intersected := make([]int64, 0)

	for _, ae := range a {
		if found := slices.Index(b, ae); found != -1 {
			intersected = append(intersected, ae)
		}
	}

	return intersected
}

// SliceStable sorts the slice data using the provided less function.
//
//	keeping equal elements in their original order.
//	It panics if x is not a slice.
func SortByKey(data []map[string]interface{}, key string, order bool) {
	// Define a custom sorting function
	sort.SliceStable(data, func(i, j int) bool {
		if _, ok := data[i][key].(string); ok {
			if order {
				return data[i][key].(string) > data[j][key].(string)
			} else {
				return data[i][key].(string) < data[j][key].(string)
			}
		}
		if _, ok := data[i][key].(int); ok {
			if !order {
				return data[i][key].(int) > data[j][key].(int)
			} else {
				return data[i][key].(int) < data[j][key].(int)
			}
		}
		return true
	})
}

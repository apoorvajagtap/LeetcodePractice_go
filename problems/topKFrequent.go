// Question link: https://leetcode.com/problems/top-k-frequent-elements/description/

// Runtime: 17 ms
// Memory: 6.9 MB

// LOGIC:
// 1. Create a map to count the occurences of each value in the sent array.
// 2. Sort the values in map, we used an additional slice for it (count).
// 3. Now that, count slice is sorted in ascending order, we found the last greatest value.  e.g, if k is 2, i.e.
// need to find 2 most frequent numbers, so, minOccurence will hold the value of 2nd greatest number.
// 4. Loop through the occurence map, and if value of any pair is >= minOccurence, append that to result slice.

package problems

import "sort"

func TopKFrequent(nums []int, k int) []int {
	occurence := make(map[int]int, len(nums))
	res := []int{}
	for _, ele := range nums {
		occurence[ele]++
	}

	// sort the 'occurence' map as per the values
	count := make([]int, len(nums))
	for _, val := range occurence {
		count = append(count, val)
	}
	sort.Ints(count)

	// find the first index that holds value less than k; e.g, if k is 2, i.e.
	// need to find 2 most frequent numbers,
	// so, minOccurence will hold the value of 2nd greatest number.
	minOccurence := count[len(count)-k]

	for key, val := range occurence {
		if val >= minOccurence {
			res = append(res, key)
		}
	}

	return res
}

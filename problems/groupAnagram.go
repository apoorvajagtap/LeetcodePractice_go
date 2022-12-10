// Question link: https://leetcode.com/problems/group-anagrams/description/

// Runtime: 37 ms
// Memory: 9.4 MB

// LOGIC:
// TLDR; just focus on 1 point below.
// - Create a map with key as array of int with size 26(for each alphabet), and value as 'slice of string'
// - Loop through each element, and increase counter of respective index (char-'a', i.e., ASCII value)
// - If same array already exists, append the element to the slice of that string.
// - Else, create new key-value pair using the new array and element.

package problems

func GroupAnagrams(strs []string) [][]string {
	dict := make(map[[26]int][]string)
	res := [][]string{}

	for _, ele := range strs {
		templs := [26]int{}
		for _, c := range ele {
			templs[c-'a'] += 1
		}

		dict[templs] = append(dict[templs], ele)
	}

	for _, v := range dict {
		res = append(res, v)
	}
	return res
}

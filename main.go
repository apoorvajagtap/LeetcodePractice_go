package main

import (
	"fmt"

	"github.com/apoorvajagtap/LeetcodePractice_go/problems"
)

func main() {
	// // call any function to test from the same package
	// in := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// out := problems.GroupAnagrams(in)

	in := []int{1, 1, 1, 2, 2, 3}
	out := problems.TopKFrequent(in, 3)

	fmt.Println(out)
}

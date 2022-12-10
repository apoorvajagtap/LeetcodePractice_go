package main

import (
	"fmt"

	"github.com/apoorvajagtap/LeetcodePractice_go/problems"
)

func main() {
	// call any function to test from the same package
	in := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	out := problems.GroupAnagrams(in)

	fmt.Println(out)
}

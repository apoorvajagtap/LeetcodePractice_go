// Question link: https://leetcode.com/problems/product-of-array-except-self/description/

// Runtime: 23 ms
// Memory: 7.6 MB

// LOGIC:
// 1. Use a counter to count number of zeroes in nums, and multiple all the rest of the elements and store in 'totalProd'.
// 2. If 0 has occurred more than once, all the elements in result list will be 0.
// 3. If 0 has occured only once, then the index at 0's place will be product of all other elements &
// rest all will be 0 (coz eventually multiplied by 0).
// 4. If 0's occurence is 0, then simply divide the totalProd with the ith element, and store the ans at same index in res.

package problems

func ProductExceptSelf(nums []int) []int {
	totalProd, countZero := 1, 0
	res := make([]int, len(nums))

	for _, ele := range nums {
		if ele != 0 {
			totalProd *= ele
		} else {
			countZero++
		}
	}

	if countZero > 1 {
		return res
	} else if countZero == 1 {
		for idx, ele := range nums {
			if ele == 0 {
				res[idx] = totalProd
			} else {
				res[idx] = 0
			}
		}
	} else {
		for idx, ele := range nums {
			res[idx] = totalProd / ele
		}
	}

	return res
}

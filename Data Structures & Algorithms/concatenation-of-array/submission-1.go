func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int,2*n)

	if n <= 0 {
		return ans
	}

	for i := range nums {
		ans[i] = nums[i]
	}

	for i := range nums {
		ans[n+i] = nums[i] 
	}

	return ans
}

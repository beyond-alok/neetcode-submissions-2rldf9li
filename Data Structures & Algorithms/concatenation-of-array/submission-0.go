func getConcatenation(nums []int) []int {
	size := 2*len(nums)
   ans := make([]int,size)
   n := len(nums) 
   
   for i := range ans {
		if i < n {
			ans[i] = nums[i]
		} else {
			ans[i] = nums[i-n]
		}

   }

   return ans

}

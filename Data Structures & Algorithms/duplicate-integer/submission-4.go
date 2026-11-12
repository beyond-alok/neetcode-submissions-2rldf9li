func hasDuplicate(nums []int) bool {
	has := make(map[int]bool)
   for _,v := range nums {
		if has[v] {
			return true
		}
		has[v] = true
   } 
   return false
}

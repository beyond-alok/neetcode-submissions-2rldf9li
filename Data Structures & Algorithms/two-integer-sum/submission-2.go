func twoSum(nums []int, target int) []int {
	var hashmap = make(map[int]int)
	for i,v := range nums {
		if n,ok := hashmap[target - v]; ok  {
			return []int{n,i}
		}
		hashmap[v] = i
	}
	return []int{}
}

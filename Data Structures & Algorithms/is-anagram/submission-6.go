func isAnagram(s string, t string) bool {
	var hashmap = make(map[rune]int)
	for _,v := range s {
		hashmap[v]++
	}

	for _,v := range t {
		if _,ok := hashmap[v]; ok {
			hashmap[v]--
		} else {
			hashmap[v]++
		}
	}

	for _,v := range hashmap{
		if v != 0 {
			return false
		} 
	}
	return true


}

func isAnagram(s string, t string) bool {
	var arr [26]int

	for _,v := range s {
		n := int(v - 'a')
		arr[n]++
	}

	for _,v := range t {
		n := int(v-'a')
		arr[n]--
	}

	for _,v := range arr {
		if v != 0 {
			return false
		}
	}
	return true

}

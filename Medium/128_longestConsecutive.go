func longestConsecutive(nums []int) int {
	var ht = make(map[int]bool)
	var res = 0

	// Initialize the map with false (not visited) and to get rid of duplicates
	for _, num := range nums {
		ht[num] = false
	}

	// iterate through the hashtable
	for num := range ht {

		// if already visited
		if ht[num] {
			continue
		}

		var currRes = 1
		var currNum = num

		// explore backward sequence
		for {
			if _, ok := ht[currNum-1]; ok {
				ht[currNum-1] = true
				currNum--
				currRes++
			} else {
				currNum = num
				break
			}
		}

		// explore forward sequence
		for {
			if _, ok := ht[currNum+1]; ok {
				ht[currNum+1] = true
				currNum++
				currRes++
			} else {
				break
			}
		}

		if currRes > res {
			res = currRes
		}
	}

	return res
}
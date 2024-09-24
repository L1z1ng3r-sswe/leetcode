package main

import "fmt"

func main() {
	fmt.Println(uniqueLetterString("ABCA")) ab,abc,abca,b,bc,bca,
}

func uniqueLetterString(s string) int {
	n := len(s)
	idx := make(map[byte][]int)

	for i := 0; i < n; i++ {
		idx[s[i]] = []int{-1, -1} // prev and next values
	}

	var res int

	for i := 0; i < n; i++ {
		ch := s[i]
		prev1 := idx[ch][0]
		prev2 := idx[ch][1]

		res += (prev1 - prev2) * (i - prev1)
		fmt.Printf(" %s %v (%d - %d) * (%d - %d) = %d \n\n", string(string(ch)), idx[ch], prev1, prev2, i, prev1, (prev1-prev2)*(i-prev1))
		fmt.Printf("res + (prev1 - prev2) * (n - prev1) = %d\n\n\n\n", res+(prev1-prev2)*(i-prev1))

		idx[ch] = []int{i, prev1}
	}

	fmt.Println(len(idx))
	for letter, arr := range idx {

		prev1 := arr[0]
		prev2 := arr[1]
		fmt.Printf(" %s %v (%d - %d) * (%d - %d) = %d \n\n", string(letter), arr, prev1, prev2, n, prev1, (prev1-prev2)*(n-prev1))
		fmt.Printf("res + (prev1 - prev2) * (n - prev1) = %d\n\n\n\n", res+(prev1-prev2)*(n-prev1))

		res += (prev1 - prev2) * (n - prev1)
	}

	return res
}

// infinity := []int{0}

// for i := 0; i < len(infinity); i++ {
// 	if i%60 == 0 {
// 		fmt.Printf("---%d---\n", i/60)
// 	} else {
// 		fmt.Println(i)
// 	}
// 	infinity = append(infinity, i+1)

// 	time.Sleep(time.Second)
// }

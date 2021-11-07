5918. 统计字符串中的元音子字符串func isTarget(a []int) bool {
	for i:=0; i<len(a); i++ {
		if a[i]==0 {
			return false
		}
	}
	return true
}

func countVowelSubstrings(word string) int {
	bytemap := map[rune]int{'a':0, 'e':1, 'i':2, 'o':3, 'u':4}
	ans := 0
	for i,ch := range word {
		if _,ok:=bytemap[ch]; !ok {
			continue
		}
		cnts := make([]int, 5)
		for j:=i; j>=0; j-- {
			if _,ok := bytemap[rune(word[j])]; !ok {
				break
			}
			cnts[bytemap[rune(word[j])]]++
			if isTarget(cnts) {
				ans++
			}
		}
	}
	return ans
}
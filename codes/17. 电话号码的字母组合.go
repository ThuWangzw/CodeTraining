//典型的排列型回溯问题
func letterCombinationsIter(digits string, num2str map[uint8]string) []string {
	results := make([]string, 0)
	if len(digits) == 0 {
		results := append(results, "")
		return results
	}
	cstr := num2str[digits[0]]
	for _, c := range cstr {
		str := string(c)
		strs := letterCombinationsIter(digits[1:], num2str)
		for _, s := range strs {
			results = append(results, str+s)
		}
	}
	return results
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}
	num2str := make(map[uint8]string)
	num2str['2'] = "abc"
	num2str['3'] = "def"
	num2str['4'] = "ghi"
	num2str['5'] = "jkl"
	num2str['6'] = "mno"
	num2str['7'] = "pqrs"
	num2str['8'] = "tuv"
	num2str['9'] = "wxyz"
	return letterCombinationsIter(digits, num2str)
}
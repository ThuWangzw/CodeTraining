type Nums []string

func (nums *Nums) Len() int {
	return len((*nums))
}

func (nums *Nums) Less(i, j int) bool {
	a := (*nums)[i] + (*nums)[j]
	b := (*nums)[j] + (*nums)[i]
	return strings.Compare(a, b) < 0
}

func (nums *Nums) Swap(i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func minNumber(nums []int) string {
	numsStr := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		numsStr[i] = strconv.Itoa(nums[i])
	}
	numstrP := Nums(numsStr)
	sort.Sort(&numstrP)
	ans := strings.Builder{}
	for _, numstr := range numsStr {
		ans.WriteString(numstr)
	}
	return ans.String()
}
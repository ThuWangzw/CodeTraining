func subsetsWithDup_iter(nums map[int]int, keys []int) [][]int {
	n := len(keys)
	if n==0 {
		result := make([][]int, 1)
		result[0] = make([]int, 0)
		return result
	}
	num := keys[0]
	cnt := nums[num]
	results := subsetsWithDup_iter(nums, keys[1:])
	m := len(results)
	for i:=1; i<=cnt; i++ {
		arys := make([]int, i)
		for j:=0; j<i; j++ {
			arys[j]=num
		}
		for j:=0; j<m; j++ {
			results = append(results, append(arys, results[j]...))
		}
	}
	return results
}

func subsetsWithDup(nums []int) [][]int {
	//init map
	num_map := make(map[int]int)
	for _, num := range(nums) {
		num_map[num]++
	}
	keys := make([]int, 0)
	for key := range(num_map) {
		keys = append(keys, key)
	}
	return subsetsWithDup_iter(num_map, keys)
}
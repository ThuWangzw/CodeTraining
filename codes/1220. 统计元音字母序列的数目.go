const (
	mod = 1000000007
)

func countVowelPermutation(n int) int {
	nextMap := map[byte][]byte{
		'a': {'e'},
		'e': {'a', 'i'},
		'i': {'a', 'e', 'o', 'u'},
		'o': {'i', 'u'},
		'u': {'a'},
	}
	counts := map[byte]int64{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}
	for i := 2; i <= n; i++ {
		newcounts := make(map[byte]int64)
		for ch, nexts := range nextMap {
			for _, next := range nexts {
				newcounts[ch] = (newcounts[ch] + counts[next]) % mod
			}
		}
		counts = newcounts
	}
	var ans int64
	for ch := range nextMap {
		ans = (ans + counts[ch]) % mod
	}
	return int(ans)
}
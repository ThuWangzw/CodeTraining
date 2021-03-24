// 简单的回溯题

func restoreIpAddressesIter(s string, ipIdx int) []string {
	n := len(s)
	results := make([]string, 0)
	for i := 0; i < 3 && i < n; i++ {
		if s[0] == '0' && i > 0 {
			break
		}
		num_str := s[:i+1]
		num, _ := strconv.Atoi(num_str)
		if num > 255 {
			break
		}
		if ipIdx == 3 && i < n-1 {
			continue
		}
		if ipIdx == 3 {
			results = append(results, num_str)
		} else {
			num_str += "."
			strs := restoreIpAddressesIter(s[i+1:], ipIdx+1)
			for _, str := range strs {
				results = append(results, num_str+str)
			}
		}
	}
	return results
}

func restoreIpAddresses(s string) []string {
	return restoreIpAddressesIter(s, 0)
}
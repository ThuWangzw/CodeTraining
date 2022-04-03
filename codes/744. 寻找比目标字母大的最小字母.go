func nextGreatestLetter(letters []byte, target byte) byte {
	ans := target
	mindis := 26
	for _, c := range letters {
		dis := int(c) - int(target)
		if dis <= 0 {
			dis += 26
		}
		if dis < mindis {
			ans = c
			mindis = dis
		}
	}
	return ans
}
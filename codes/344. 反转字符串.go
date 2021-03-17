func reverseString(s []byte)  {
	m := len(s)
	n := len(s)/2
	for i:=0; i<n; i++ {
		tmp := s[i]
		s[i] = s[m-1-i]
		s[m-1-i] = tmp
	}
}
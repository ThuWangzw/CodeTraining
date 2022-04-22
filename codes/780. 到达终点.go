func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for tx >= sx && ty >= sy {
		if sx == tx && sy == ty {
			return true
		}
		if tx == ty {
			break
		} else if tx > ty {
			k := max(1, (tx-sx)/ty)
			tx -= k * ty
		} else {
			k := max(1, (ty-sy)/tx)
			ty -= k * tx
		}
	}
	return false
}
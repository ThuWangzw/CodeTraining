type StockSpanner struct {
	stocks []int
	stack  []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stocks: make([]int, 0),
		stack:  make([]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	this.stocks = append(this.stocks, price)
	for len(this.stack) > 0 && this.stocks[this.stack[len(this.stack)-1]] <= price {
		this.stack = this.stack[:len(this.stack)-1]
	}
	ans := len(this.stocks)
	if len(this.stack) > 0 {
		ans = len(this.stocks) - 1 - this.stack[len(this.stack)-1]
	}
	this.stack = append(this.stack, len(this.stocks)-1)
	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
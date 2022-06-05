type TextEditor struct {
	text   []byte
	cursor int
}

func Constructor() TextEditor {
	return TextEditor{
		text:   make([]byte, 0),
		cursor: 0,
	}
}

func (this *TextEditor) AddText(text string) {
	cursorRight := []byte(text)
	cursorRight = append(cursorRight, this.text[this.cursor:]...)
	this.text = append(this.text[:this.cursor], cursorRight...)
	this.cursor += len(text)
}

func (this *TextEditor) DeleteText(k int) int {
	k = min(k, this.cursor)
	this.text = append(this.text[:this.cursor-k], this.text[this.cursor:]...)
	this.cursor -= k
	return k
}

func (this *TextEditor) CursorLeft(k int) string {
	k = min(k, this.cursor)
	this.cursor -= k
	begin := max(0, this.cursor-10)
	return string(this.text[begin:this.cursor])
}

func (this *TextEditor) CursorRight(k int) string {
	this.cursor += k
	if this.cursor > len(this.text) {
		this.cursor = len(this.text)
	}
	begin := max(0, this.cursor-10)
	return string(this.text[begin:this.cursor])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
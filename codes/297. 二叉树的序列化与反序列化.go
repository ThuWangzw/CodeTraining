/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize_it(root *TreeNode, builder *strings.Builder) {
	if root == nil {
		if builder.Len() == 0 {
			builder.WriteString("n")
		} else {
			builder.WriteString(",n")
		}
		return
	}
	if builder.Len() == 0 {
		builder.WriteString(strconv.Itoa(root.Val))
	} else {
		builder.WriteString("," + strconv.Itoa(root.Val))
	}
	this.serialize_it(root.Left, builder)
	this.serialize_it(root.Right, builder)
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	builder := &strings.Builder{}
	this.serialize_it(root, builder)
	return builder.String()
}

func (this *Codec) deserialize_it(data *string, begin int) (*TreeNode, int) {
	var root *TreeNode
	i := begin
	for ; i < len((*data)); i++ {
		if (*data)[i] == ',' {
			break
		}
	}
	num, err := strconv.Atoi((*data)[begin:i])
	if err != nil {
		return nil, i + 1
	}
	root = &TreeNode{
		Val: num,
	}
	if i != len((*data)) {
		Left, rightbegin := this.deserialize_it(data, i+1)
		Right, nextbegin := this.deserialize_it(data, rightbegin)
		root.Left = Left
		root.Right = Right
		return root, nextbegin
	} else {
		return root, len((*data))
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	root, _ := this.deserialize_it(&data, 0)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
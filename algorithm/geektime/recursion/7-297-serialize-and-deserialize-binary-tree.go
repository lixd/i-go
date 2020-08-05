package recursion

// https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
// TODO 二叉树序列化与反序列化
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return ""

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	return nil
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

package recursion

// https://leetcode-cn.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	var res []string
	generate(0, 0, n, "", &res)
	return res
}
func generate(left, right, n int, s string, res *[]string) {
	// terminator
	if left == n && right == n {
		*res = append(*res, s)
		return
	}
	// process current logic
	// drill down
	if left < n {
		generate(left+1, right, n, s+"(", res)
	}
	if left > right {
		generate(left, right+1, n, s+")", res)
	}
	// restore current status
}

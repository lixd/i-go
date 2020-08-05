package recursion

// recursion 递归 模板代码
func recursion(level, params int) {
	// 1.首先写递归终结条件 terminator
	if level > 999 {
		// process result
		return
	}
	// 2.处理当前层逻辑 process current logic
	process(level, params)
	// 3.下探到下一层 drill down
	recursion(level+1, params)
	// 4.清理当前层 restore current status
}

func process(level, params int) {

}

package bfsdfs

// 深度优先 每次遍历直接进入下一层循环,如果有的话,直到某个分支循环到底了再返回到最开始的循环
/*
  visited =set()

  def dfs(node,visited):
    // 如果访问过 直接返回
	if node in visited {
		return
	}
	// 否则添加到已访问列表
    visited.add(node)
	// process current node here.
	for next_node in node.children()
       // // 不等当前层循环走完 直接进入下一层循环
	  if not next_ndoe in visited:
        dfs(next_node,visited)
*/

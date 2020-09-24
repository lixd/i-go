package bfsdfs

// 广度优先 每次遍历时记录当前层的结果并将下一层节点存储起来等待下一次遍历
/*
  visited = set()
  def bfs(graph,start,end):

   queue = []
   queue.append([start])
   visited.add(start)

   while queue:
     node=queue.pop()
     visited.add(node)
     process(node)
     // 每次存储当前节点相关联的下一层级节点
     nodes = generate_related_nodes(node)
     queue.push(nodes)
// other process working
*/

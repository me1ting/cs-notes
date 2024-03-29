# 最短路径

**最短路径问题**指求出图中两个顶点之间，所有路径的最小值。

最短路径问题受到以下因素的影响：

- 从固定顶点出发，还是任意定点出发？分别用**单源**或**所有点对**来修饰
- 边是否具有权重，分别用**有权**或**无权**修饰
- 有权的图，其边是否允许负值，用**负边值的图**修饰
- ...

# 单源最短路径

从固定顶点出发，计算其到任意其它定点的最短路径。

## 单源无权最短路径

使用**广度优先搜索算法**可以完成计算。

```go
// 假设顶点序号从1开始，使用邻接表表示，源点为1
func search(G [][]int)[]int{
	n:=len(G)
	//使用一个数组记录最短路径
	paths :=make([]int,n)
	for i:=0;i<n;i++ {
		path[i] = -1//使用负数表示路径未知
	}
	paths[0] = 0

	//使用一个队列记录按层遍历的顶点
	queue = []int{0}
	for len(queue) > 0 {
		v:=queue[0]
		for _,adj := range G[v] {
			if paths[adj] == -1 {
				paths[adj] = paths[v]+1
				queue = append(queue, adj)
			}
		}
		queue = queue[1:]
	}
	return paths
}
```

## 单源有权最短路径

使用**Dijkstra**算法来完成，其基本思想是：

- 每次从未访问但邻接已访问顶点的顶点中，选出已知最短路径最小的顶点，将该顶点标记为已访问，并更新其邻接未访问节点的已知最短路径
- 被选出的顶点因为已知最短路径最小  ，不存在经过未访问节点的更小路径，因此可以确定其最小路径就是已知最小路径

//TODO: add code

# 参考资料

《数据结构与算法分析 Java语言描述》v3 9.3 最短路径
《算法》v4 4.4 最短路径
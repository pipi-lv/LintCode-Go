package leetcode

import "container/heap"

// Node 定义一个存储节点和离起点相应距离的数据结构
type Node struct {
	Val  int
	Cost int
}

type Queue []*Node

func (pq *Queue) Len() int { return len(*pq) }

func (pq *Queue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return (*pq)[i].Cost < (*pq)[j].Cost
}

func (pq *Queue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *Queue) Push(x any) {
	*pq = append(*pq, x.(*Node))
}

func (pq *Queue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func minimumTime(n int, edges [][]int, disappear []int) []int {
	paths := make([][]struct{ v, length int }, n)
	for _, edge := range edges {
		u, v, length := edge[0], edge[1], edge[2]
		paths[u] = append(paths[u], struct{ v, length int }{v, length})
		paths[v] = append(paths[v], struct{ v, length int }{u, length})
	}

	roads := make([]int, n)
	for i := 1; i < n; i++ {
		roads[i] = -1
	}

	pq := make(Queue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Node{})

	for pq.Len() > 0 {
		t := heap.Pop(&pq).(*Node)
		if t.Cost != roads[t.Val] {
			continue
		}

		for _, edge := range paths[t.Val] {
			v, length := edge.v, edge.length
			if t.Cost+length < disappear[v] && (roads[v] == -1 || t.Cost+length < roads[v]) {
				heap.Push(&pq, &Node{v, t.Cost + length})
				roads[v] = t.Cost + length
			}
		}
	}

	return roads
}

// 堆优化
func minimumTime2(n int, edges [][]int, disappear []int) []int {
	// 构造邻接矩阵 TODO:这边数组强行初始化太耗空间，改为结构体且动态append
	paths := make([][]int, n)
	for i := 0; i < n; i++ {
		paths[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				paths[i][j] = 100001
			}
		}
	}

	for _, edge := range edges {
		a, b, c := edge[0], edge[1], edge[2]
		if a != b {
			if c < disappear[b] {
				paths[a][b] = min(paths[a][b], c)
			}

			if c < disappear[a] {
				paths[b][a] = min(paths[b][a], c)
			}
		}
	}

	// 改用Dijkstra算法
	roads := make([]int, n)
	for i := 1; i < n; i++ {
		roads[i] = -1
	}

	// 初始化队列、未访问节点
	pq := make(Queue, 0)
	heap.Init(&pq)

	// 将起点加入pq
	heap.Push(&pq, &Node{})

	for pq.Len() > 0 {
		t := heap.Pop(&pq).(*Node)
		// 当前节点是终点，即可返回最短路径
		if t.Cost != roads[t.Val] {
			continue
		}

		for index, length := range paths[t.Val] {
			if t.Cost+length < disappear[index] && (roads[index] == -1 || t.Cost+length < roads[index]) {
				heap.Push(&pq, &Node{index, t.Cost + length})
				roads[index] = t.Cost + length
			}
		}
	}

	return roads
}

// 堆优化
func minimumTime3(n int, edges [][]int, disappear []int) []int {
	// 构造邻接矩阵
	paths := make([][]int, n)
	for i := 0; i < n; i++ {
		paths[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				paths[i][j] = 100001
			}
		}
	}

	for _, edge := range edges {
		a, b, c := edge[0], edge[1], edge[2]
		if a != b {
			if c < disappear[b] {
				paths[a][b] = min(paths[a][b], c)
			}

			if c < disappear[a] {
				paths[b][a] = min(paths[b][a], c)
			}
		}
	}

	// 改用Dijkstra算法
	roads := make([]int, n)
	for i := 1; i < n; i++ {
		roads[i] = -1
	}

l:
	for i := 1; i < n; i++ {
		visited := make([]bool, n)

		// 初始化队列、未访问节点
		pq := make(Queue, 0)
		heap.Init(&pq)

		// 将起点加入pq
		heap.Push(&pq, &Node{})
		for pq.Len() > 0 {
			t := heap.Pop(&pq).(*Node)
			// 当前节点是终点，即可返回最短路径
			if t.Val == i {
				if t.Cost < disappear[i] {
					roads[i] = t.Cost
				}
				continue l
			}

			// 若当前节点已遍历过，跳过当前节点
			if visited[t.Val] {
				continue
			}

			// 将当前节点标记成已遍历
			visited[t.Val] = true
			for j := 0; j < n; j++ {
				if paths[t.Val][j] != 100001 && !visited[j] && t.Cost+paths[t.Val][j] < disappear[j] {
					heap.Push(&pq, &Node{Val: j, Cost: t.Cost + paths[t.Val][j]})
				}
			}
		}
	}

	return roads
}

// 第一次提交未考虑中间点到时间消失就不能通过的问题
// 第二次提交未考虑自身到自身点最小为0，不需要考虑传入的数组edges旳值 ✓
// 第三次提交未考虑断路问题，应把所有值初始化为100001 ✓
// 第四次提交未考虑无向图问题，路径顺序问题 ✓
// 第五次提交未考虑相同两个点最短路径覆盖问题 ✓
// 第六次提交核心在于计算最短路径时要考虑到消失时间问题，和第一次问题相同但是要考虑消失时间顺序累加问题，从某些点到另一个点的距离可能小于消失时间，但从0出发时候可能就大于了
// 综上所述这题应该用Dijkstra而不是弗洛伊德算法，算法过程更简单，弗洛伊德算法判断过程很复杂甚至无法判断
// 第七次超时，需要用堆来优化Dijkstra算法
func minimumTime4(n int, edges [][]int, disappear []int) []int {
	// 构造邻接矩阵
	paths := make([][]int, n)
	for i := 0; i < n; i++ {
		paths[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i != j {
				paths[i][j] = 100001
			}
		}
	}

	for _, edge := range edges {
		a, b, c := edge[0], edge[1], edge[2]
		if a != b {
			if c < disappear[b] {
				paths[a][b] = min(paths[a][b], c)
			}

			if c < disappear[a] {
				paths[b][a] = min(paths[b][a], c)
			}
		}
	}

	// 改用Dijkstra算法
	roads, visited := make([]int, n), make([]bool, n)

	// 初始化
	roads[0] = 0
	visited[0] = true
	for i := 1; i < n; i++ {
		roads[i] = -1
	}

	for i := 1; i < n; i++ {
		// 距离 start 最近的点
		k := -1
		// 距离 start 最近的距离
		disMin := 100001
		// 1. 选取出距离顶点 start 最近的一个顶点
		for j := 1; j < n; j++ {
			// 元素不在已访问的列表中且
			if visited[j] == false && paths[0][j] < disMin {
				disMin = paths[0][j]
				k = j
			}
		}

		if k > 0 {
			// 更新信息,加入到最短的集合
			visited[k] = true
			if roads[i] < disappear[i] {
				roads[k] = disMin
			}

			// 更新距离表
			for index := 1; index < n; index++ {
				// 1. 不在最短列表中
				// 2. start==>shortestIndex+si==>index < start==>index，则进行距离表更新
				if visited[index] == false && paths[0][k]+paths[k][index] < paths[0][index] && paths[0][k] < disappear[k] && paths[0][k]+paths[k][index] < disappear[index] {
					paths[0][index] = paths[0][k] + paths[k][index]
				}
			}
		}
	}

	return roads
}

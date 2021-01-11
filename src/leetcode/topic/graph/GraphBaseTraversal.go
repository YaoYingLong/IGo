package graph

import (
	_ "leetcode/topic"
	"sort"
)

/*
 1030
 距离顺序排列矩阵单元格
 思路：
 - 由于输出是一个2列 R* C行的二维数组
 - 所以先将矩阵转换为上面说得二维数组
 - 再对二维数组中的数据排序
*/
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	ret := make([][]int, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			ret = append(ret, []int{i, j})
		}
	}

	sort.Slice(ret, func(i, j int) bool {
		a, b := ret[i], ret[j]
		return abs(a[0]-r0)+abs(a[1]-c0) < abs(b[0]-r0)+abs(b[1]-c0)
	})
	return ret
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*
 * 797
 * 所有可能的路径
 * 输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
 * 输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
 * <p>
 * 题目中给出的图是有向无环的，那么我们可以通过深度优先搜索的方法，递归求解出所有的路径。
 * <p>
 * 设图中有 N 个节点，在搜索时，如果我们到达了节点 N - 1，那么此时的路径就为 {N - 1}；
 * 否则如果我们到达了其它的节点 node，那么路径就为 {node} 加上 {所有从 nei 到 N - 1} 的路径集合，其中 nei 是 node 直接相邻的节点。
 */
func AllPathsSourceTarget(graph [][]int) [][]int {
	return allPathsSourceTargetDfs(graph, 0)
}

func allPathsSourceTargetDfs(graph [][]int, node int) [][]int {
	var ans [][]int
	N := len(graph)
	if node == N-1 {
		path := []int{N - 1}
		ans = append(ans, path)
		return ans
	}

	for _, val := range graph[node] {
		for _, paths := range allPathsSourceTargetDfs(graph, val) {
			paths = append([]int{node}, paths...)
			ans = append(ans, paths)
		}
	}
	return ans
}

package combination

/*
	给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
	示例:

	输入: n = 4, k = 2
	输出:
	[
	[2,4],
	[3,4],
	[2,3],
	[1,2],
	[1,3],
	[1,4],
	]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/combinations
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func combine(n int, k int) [][]int {
	res := [][]int{}
	tmp := []int{}
	res = dfs(res, tmp, n, k, 0)
	return res
}

func dfs(res [][]int, tmp []int, n, k, x int) [][]int {
	if k == len(tmp) {
		tmpCopy := make([]int, len(tmp))
		copy(tmpCopy, tmp)
		res = append(res, tmpCopy)
		return res
	}
	for i := x + 1; i <= n; i++ {
		tmp = append(tmp, i)
		res = dfs(res, tmp, n, k, i)
		tmp = tmp[:len(tmp)-1]
	}

	return res
}
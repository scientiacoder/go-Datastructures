package permutation

/*
	给定一个 没有重复 数字的序列，返回其所有可能的全排列。
	示例:

	输入: [1,2,3]
	输出:
	[
	[1,2,3],
	[1,3,2],
	[2,1,3],
	[2,3,1],
	[3,1,2],
	[3,2,1]
	]

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/permutations
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	visited := make([]bool, len(nums))
	res := [][]int{}
	tmp := []int{}
	res = dfs(nums, visited, res, tmp, len(nums), 0)
	return res
}

func dfs(nums []int, visited []bool, res [][]int, tmp []int, length, depth int) [][]int {
	if length == depth {
		tmpCopy := make([]int, len(tmp))
		copy(tmpCopy, tmp)
		res = append(res, tmpCopy)
		return res
	}
	for i, num := range nums {
		if visited[i] == false {
			tmp = append(tmp, num)
			visited[i] = true
			res = dfs(nums, visited, res, tmp, length, depth+1)
			//fmt.Println(tmp)
			visited[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
	return res
}

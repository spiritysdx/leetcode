// https://leetcode.cn/problems/step-by-step-directions-from-a-binary-tree-node-to-another/
func getDirections(root *TreeNode, startValue, destValue int) string {
	// 定义一个字节数组，用于存储路径
	var path []byte
	// 定义深度优先搜索函数
	var dfs func(*TreeNode, int) bool
	dfs = func(node *TreeNode, target int) bool {
		if node == nil { // 如果当前节点为空，返回false
			return false
		}
		if node.Val == target { // 如果当前节点的值等于目标值，返回true
			return true
		}
		path = append(path, 'L') // 尝试向左子树搜索
		if dfs(node.Left, target) {
			return true // 如果找到目标值，返回true
		}
		path[len(path)-1] = 'R' // 回溯并尝试向右子树搜索
		if dfs(node.Right, target) {
			return true // 如果找到目标值，返回true
		}
		path = path[:len(path)-1] // 回溯，移除路径的最后一个方向
		return false
	}
	// 搜索从根节点到起始值的路径
	dfs(root, startValue)
	pathToStart := path // 保存从根节点到起始值的路径
	path = nil // 清空路径变量
	// 搜索从根节点到目标值的路径
	dfs(root, destValue)
	pathToDest := path // 保存从根节点到目标值的路径
	// 找到起始路径和目标路径的公共前缀，并移除
	for len(pathToStart) > 0 && len(pathToDest) > 0 && pathToStart[0] == pathToDest[0] {
		pathToStart = pathToStart[1:] // 去掉路径中公共的部分
		pathToDest = pathToDest[1:]
	}
	// 构造最终的路径：从起始值到目标值的路径
	// 先向上移动若干步（路径中 "U" 的个数），再按照目标路径的方向移动
	return strings.Repeat("U", len(pathToStart)) + string(pathToDest)
}

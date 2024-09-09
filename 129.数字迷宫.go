type pair struct{ x, y int } // 定义一个结构体，表示网格中的方向对 (x, y)

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右方向

// wordPuzzle 函数检查单词是否存在于网格中
func wordPuzzle(grid [][]byte, word string) bool {
	h, w := len(grid), len(grid[0]) // 获取网格的高度和宽度
	vis := make([][]bool, h) // 创建一个与网格相同大小的布尔矩阵，用于标记访问状态
	for i := range vis {
		vis[i] = make([]bool, w) // 初始化每行的布尔矩阵
	}
	var exist func(i, j, k int) bool // 定义一个递归函数用于检查单词是否存在
	exist = func(i, j, k int) bool {
		if grid[i][j] != word[k] { // 剪枝：当前字符不匹配，直接返回 false
			return false
		}
		if k == len(word)-1 { // 如果已经匹配到单词的最后一个字符，则说明找到了该单词
			return true
		}
		vis[i][j] = true // 标记当前单元格为已访问
		defer func() { vis[i][j] = false }() // 回溯时还原单元格的访问状态
		for _, dir := range directions { // 遍历所有四个方向
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] { // 检查新坐标是否在网格内且未被访问
				if exist(newI, newJ, k+1) { // 递归检查新位置的下一个字符
					return true
				}
			}
		}
		return false // 所有方向都检查完后，返回 false
	}
	for i, row := range grid { // 遍历网格的每个单元格
		for j := range row {
			if exist(i, j, 0) { // 从每个单元格开始检查是否能找到单词
				return true
			}
		}
	}
	return false // 遍历完所有单元格后，若未找到单词，则返回 false
}

// 思路
// 递归搜索四个方向，约束下一个位置在实际矩阵内
// 遍历每一个位置，因为每一个位置都有可能是文本的起始点
// 访问过的点需要备注已访问，访问完毕需要重置状态 - 这一步可以在回溯时改变原始矩阵做到，访问过就将原始矩阵对应位置赋值为不存在的一个值

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathTarget(root *TreeNode, target int) [][]int {
    res := [][]int{} // 存储符合条件的路径
    path := []int{}  // 当前路径
    // 深度优先搜索函数
    var find func(r *TreeNode, t int)
    find = func(r *TreeNode, t int) {
        if r == nil {
            return
        }
        // 将当前节点值加入路径
        path = append(path, r.Val)
        t -= r.Val // 更新剩余目标值
        // 如果是叶子节点并且路径和等于目标值，保存该路径
        if r.Left == nil && r.Right == nil && t == 0 {
            res = append(res, append([]int{}, path...)) // 深拷贝当前路径
        }
        // 递归遍历左右子树
        find(r.Left, t)
        find(r.Right, t)
        // 回溯，移除当前节点值，也就是最后一个节点
        path = path[:len(path)-1]
    }
    find(root, target)
    return res
}

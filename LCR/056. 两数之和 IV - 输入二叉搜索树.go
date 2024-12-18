/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    // 中序遍历获取非递减序列的切片
    res := []*TreeNode{}
    var find func(*TreeNode)
    find = func(node *TreeNode) {
        if node != nil {
            find(node.Left)
            res = append(res, node)
            find(node.Right)
        }
    }
    find(root)
    // 二分法找差值等于k的
    size := len(res)
    for index, a := range res {
        bias := k - a.Val
        left := index+1
        right := size
        for left < right {
            mid := left + (right-left)/2
            if res[mid].Val < bias {
                left = mid +1 
            } else {
                right = mid
            }
        }
        if left <= size-1 && res[left].Val == bias {
            return true
        }
    }
    return false
}

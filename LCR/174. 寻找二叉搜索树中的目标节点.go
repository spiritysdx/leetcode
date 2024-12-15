/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTargetNode(root *TreeNode, cnt int) int {
    // 创建一个数组用来存储节点值，按照中序遍历倒序顺序存储
    count := make([]int, 0)
    // 调用辅助函数进行中序遍历倒序，结果存储到数组中
    dfskthLargest(root, &count)
    // 返回数组中第 cnt 大的值（对应索引 cnt-1）
    return count[cnt-1]
}

func dfskthLargest(root *TreeNode, count *[]int) {
    // 如果右子树存在，优先遍历右子树（中序遍历的倒序）
    if root.Right != nil {
        dfskthLargest(root.Right, count)
    }
    // 访问当前节点，将节点值加入结果数组
    if root != nil {
        *count = append(*count, root.Val)
    }
    // 如果左子树存在，最后遍历左子树
    if root.Left != nil {
        dfskthLargest(root.Left, count)
    }
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    nodes := []*TreeNode{}
    inorderTraversal(root, &nodes) // 中序遍历获取所有节点
    // 重构成单调递增的树
    for i := 0; i < len(nodes)-1; i++ {
        nodes[i].Left = nil          // 清除左子节点
        nodes[i].Right = nodes[i+1] // 连接右子节点
    }
    if len(nodes) > 0 {
        nodes[len(nodes)-1].Left = nil
        nodes[len(nodes)-1].Right = nil // 最后一个节点清空左右指针
    }
    return nodes[0]
}

func inorderTraversal(root *TreeNode, nodes *[]*TreeNode) {
    if root == nil {
        return
    }
    inorderTraversal(root.Left, nodes)  // 访问左子树
    *nodes = append(*nodes, root)       // 保存当前节点
    inorderTraversal(root.Right, nodes) // 访问右子树
}

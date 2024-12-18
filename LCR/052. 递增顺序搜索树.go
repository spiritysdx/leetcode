// /**
//  * Definition for a binary tree node.
//  * type TreeNode struct {
//  *     Val int
//  *     Left *TreeNode
//  *     Right *TreeNode
//  * }
//  */
// func increasingBST(root *TreeNode) *TreeNode {
//     nodes := []*TreeNode{}
//     inorderTraversal(root, &nodes) // 中序遍历获取所有节点
//     // 重构成单调递增的树
//     for i := 0; i < len(nodes)-1; i++ {
//         nodes[i].Left = nil          // 清除左子节点
//         nodes[i].Right = nodes[i+1] // 连接右子节点
//     }
//     if len(nodes) > 0 {
//         nodes[len(nodes)-1].Left = nil
//         nodes[len(nodes)-1].Right = nil // 最后一个节点清空左右指针
//     }
//     return nodes[0]
// }

// func inorderTraversal(root *TreeNode, nodes *[]*TreeNode) {
//     if root == nil {
//         return
//     }
//     inorderTraversal(root.Left, nodes)  // 访问左子树
//     *nodes = append(*nodes, root)       // 保存当前节点
//     inorderTraversal(root.Right, nodes) // 访问右子树
// }

func increasingBST(root *TreeNode) *TreeNode {
    vals := []int{} // 用于存储中序遍历的值
    var inorder func(*TreeNode)
    // 定义中序遍历函数
    inorder = func(node *TreeNode) {
        if node != nil { // 如果当前节点不为空
            inorder(node.Left)              // 递归遍历左子树
            vals = append(vals, node.Val)   // 将当前节点的值加入到切片中
            inorder(node.Right)             // 递归遍历右子树
        }
    }
    inorder(root) // 对整个树进行中序遍历

    dummyNode := &TreeNode{} // 创建一个虚拟头节点，方便构造结果树
    curNode := dummyNode     // 用于迭代构造新树的指针
    for _, val := range vals { // 遍历中序遍历结果的值
        curNode.Right = &TreeNode{Val: val} // 创建新节点并作为当前节点的右子节点
        curNode = curNode.Right             // 更新当前节点为新创建的节点
    }
    return dummyNode.Right // 返回真正的树根节点（虚拟头节点的右子节点）
}

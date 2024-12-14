/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func decorateRecord(root *TreeNode) (ans [][]int) {
    // 如果根节点为空，直接返回空结果
    if root == nil {
        return
    }
    // 初始化当前层的节点列表，初始值为根节点
    cur := []*TreeNode{root}
    // 开始层序遍历
    for len(cur) > 0 {
        nxt := []*TreeNode{} // 下一层的节点列表
        vals := make([]int, len(cur)) // 预分配空间，用于存储当前层节点的值
        // 遍历当前层的所有节点
        for i, node := range cur {
            vals[i] = node.Val // 记录当前节点的值
            // 如果左子节点存在，加入下一层节点列表
            if node.Left != nil {
                nxt = append(nxt, node.Left)
            }
            // 如果右子节点存在，加入下一层节点列表
            if node.Right != nil {
                nxt = append(nxt, node.Right)
            }
        }
        // 更新当前层为下一层
        cur = nxt
        // 将当前层的值加入结果集
        ans = append(ans, vals)
    }
    // 返回结果
    return
}

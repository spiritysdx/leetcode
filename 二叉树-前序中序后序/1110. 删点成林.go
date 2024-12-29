// https://leetcode.cn/problems/delete-nodes-and-return-forest/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 第一种思路：在它的父节点被删除，而自己没有被删除时，将自己加入答案中
// 第二种思路：在一个节点自己将被删除，但是子节点不需要被删除时，将子节点加入答案中
// 要先保证整棵树能够遍历完全，只能把处理代码放在递归后面，所以会是后序遍历
func delNodes(root *TreeNode, toDelete []int) (ans []*TreeNode) {
    set := make(map[int]struct{}, len(toDelete))
    for _, x := range toDelete {
        set[x] = struct{}{}
    }
    var dfs func(*TreeNode) *TreeNode
    dfs = func(node *TreeNode) *TreeNode {
        if node == nil {
            return nil
        }
        node.Left = dfs(node.Left)
        node.Right = dfs(node.Right)
        if _, ok := set[node.Val]; !ok {
            return node
        }
        if node.Left != nil {
            ans = append(ans, node.Left)
        }
        if node.Right != nil {
            ans = append(ans, node.Right)
        }
        return nil
    }
    if dfs(root) != nil {
        ans = append(ans, root)
    }
    return
}

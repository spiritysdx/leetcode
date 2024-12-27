// https://leetcode.cn/problems/merge-two-binary-trees/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 == nil {
        return nil
    }
    if root1 == nil && root2 != nil {
        return root2
    }
    if root1 != nil && root2 == nil {
        return root1
    }
    ans := &TreeNode{}
    ans.Left = mergeTrees(root1.Left, root2.Left)
    ans.Right = mergeTrees(root1.Right, root2.Right)
    ans.Val = root1.Val + root2.Val
    return ans
}

// func mergeTrees(root1, root2 *TreeNode) *TreeNode {
//     if root1 == nil {
//         return root2
//     }
//     if root2 == nil {
//         return root1
//     }
//     return &TreeNode{root1.Val + root2.Val,
//         mergeTrees(root1.Left, root2.Left),   // 合并左子树
//         mergeTrees(root1.Right, root2.Right)} // 合并右子树
// }
/*
如果 root1​ 是空的，无需合并，直接返回 root2​。
如果 root2​ 是空的，无需合并，直接返回 root1​。
如果都不为空，那么将这两个节点的值相加，作为合并后节点的新值。然后递归合并 root1​.left 与 root2​.left，得到合并后的左子树；递归合并 root1​.right 与 root2​.right，得到合并后的右子树。最后返回合并后的节点。
*/

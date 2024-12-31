// https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/description/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthLargestLevelSum(root *TreeNode, k int) int64 {
    cur := []*TreeNode{root}
    maxSum := []int{}
    var tempSum int
    for len(cur) > 0 {
        nxt := []*TreeNode{}
        tempSum = 0
        for _, c := range cur {
            tempSum += c.Val
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
        }
        cur = nxt
        maxSum = append(maxSum, tempSum)
    }
    size := len(maxSum)
    if size < k {
        return int64(-1)
    } else {
        slices.Sort(maxSum)
        return int64(maxSum[size-k])
    }
}

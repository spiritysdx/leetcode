// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    ans := 0
    cur := []*TreeNode{root}
    for len(cur) > 0 {
        temp := []int{}
        nxt := []*TreeNode{}
        for _, c := range cur {
            temp = append(temp, c.Val)
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
        }
        cur = nxt
        ans += 1
    }
    return ans
}

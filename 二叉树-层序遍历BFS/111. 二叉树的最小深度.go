// https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
func minDepth(root *TreeNode) int {
    ans := 0
    if root == nil {
        return ans
    }
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
            if c.Left == nil && c.Right == nil { // 在叶子节点时返回深度+1，首次出现叶子节点代表最小的树深度
                return ans+1
            }
        }
        cur = nxt
        ans += 1
    }
    return -1
}

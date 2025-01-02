// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
        return nil
    }
	cur := []*Node{root}
    for len(cur) > 0 {
        nxt := []*Node{}
        if len(cur) > 1 {
            for i := len(cur)-2; i >=0; i-- {
                cur[i].Next = cur[i+1]
            }
        }
        for _, c := range cur {
            if c.Left != nil {
                nxt = append(nxt, c.Left)
            }
            if c.Right != nil {
                nxt = append(nxt, c.Right)
            }
        }
        cur = nxt
    }
    return root
}

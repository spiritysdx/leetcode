// https://leetcode.cn/problems/swap-nodes-in-pairs/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    cur := head
    n := 0
    for cur != nil {
        n++
        cur = cur.Next
    }
    dummyNode := &ListNode{Next: head}
    p0 := dummyNode
    cur = p0.Next
    var pre *ListNode
    for n >= 2 {
        n-=2
        for i:=0; i<2; i++ {
            next := cur.Next
            cur.Next = pre
            pre = cur
            cur = next
        }
        next := p0.Next         // 暂存当前组的起点
        p0.Next.Next = cur      // 反转后的最后一个节点指向下一部分
        p0.Next = pre           // p0指向反转后的头节点
        p0 = next               // 移动到下一组的起点
    }
    return dummyNode.Next
}

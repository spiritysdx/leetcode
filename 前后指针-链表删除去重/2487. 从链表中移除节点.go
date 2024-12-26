// https://leetcode.cn/problems/remove-nodes-from-linked-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    haed2 := reverse(head)
    cur := haed2
    for cur != nil && cur.Next != nil {
        val := cur.Val
        if cur.Next.Val < val {
            for cur.Next != nil && cur.Next.Val < val {
                cur.Next = cur.Next.Next
            }
        } else {
            cur = cur.Next
        }
    }
    return reverse(haed2)
}

func reverse(head *ListNode) *ListNode {
    cur := head
    var pre *ListNode
    for cur != nil {
        nxt := cur.Next
        cur.Next = pre
        pre = cur
        cur = nxt
    }
    return pre
}

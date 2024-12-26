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

// 通过 206. 反转链表，我们可以从反转后的链表头开始，像 83. 删除排序链表中的重复元素 那样，删除比当前节点值小的元素。最后再次反转链表，即为答案。

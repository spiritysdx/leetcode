// https://leetcode.cn/problems/double-a-number-represented-as-a-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func doubleIt(head *ListNode) *ListNode {
    l, n := reverse(head)
    carry := 0 // 进位
    dummy := &ListNode{Next: l}
    cur := dummy.Next
    for i := 0; i < n; i++ {
        sum := carry
        sum += 2*cur.Val
        carry = sum/10
        cur.Val = sum%10
        if cur.Next != nil {
            cur = cur.Next
        } else if carry != 0 {
            cur.Next = &ListNode{Val: carry}
        }
    }
    res, _ := reverse(dummy.Next)
    return res
}

func reverse(L *ListNode) (*ListNode, int) {
    cur := L
    var pre *ListNode
    n := 0
    for cur != nil {
        nxt := cur.Next
        cur.Next = pre
        pre = cur
        cur = nxt
        n++
    }
    return pre, n
}

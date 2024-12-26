// https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    mid := findMiddle(head)
    head2 := reverse(mid)
    maxNum := 0
    for head != nil && head2 != nil {
        maxNum = max(maxNum, head.Val+head2.Val)
        head = head.Next
        head2 = head2.Next
    }
    return maxNum
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func findMiddle(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
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

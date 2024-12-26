// https://leetcode.cn/problems/palindrome-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 找链表中间节点，后续链表反转，然后一起遍历，若不相同就不回文，否则回文
func isPalindrome(head *ListNode) bool {
    mid := findMiddle(head)
    head2 := reverse(mid)
    for head != nil && head2 != nil {
        if head.Val != head2.Val {
            return false
        }
        head = head.Next
        head2 = head2.Next
    }
    return true
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

// 或者使用栈做或者切片做

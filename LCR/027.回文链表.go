/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil {
        return true
    }

    // 找到链表的中间位置
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // 反转链表的后半部分
    rev := reverse(slow)

    // 比较前半部分和反转后的后半部分
    for rev != nil {
        if head.Val != rev.Val {
            return false
        }
        head = head.Next
        rev = rev.Next
    }
    
    return true
}

func reverse(head *ListNode) *ListNode {
    var prev *ListNode
    for head != nil {
        next := head.Next
        head.Next = prev
        prev = head
        head = next
    }
    return prev
}

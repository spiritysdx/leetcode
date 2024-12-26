// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    right := dummy
    for i := 0; i<n; i++ {
        right = right.Next
    }
    left := dummy
    for right.Next!= nil { // left到倒数n+1个节点了
        left = left.Next
        right = right.Next
    }
    left.Next = left.Next.Next
    return dummy.Next
}

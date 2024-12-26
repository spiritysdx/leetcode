// https://leetcode.cn/problems/linked-list-cycle-ii/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast { // 如果有环，那么从头节点出发的指针和慢指针继续走，会在入口节点处相遇
            for slow != head {
                slow = slow.Next
                head = head.Next
            }
            return slow
        }
    }
    return nil
}

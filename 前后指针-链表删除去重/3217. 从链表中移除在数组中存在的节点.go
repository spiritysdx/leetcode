// https://leetcode.cn/problems/delete-nodes-from-linked-list-present-in-array/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    removeMap := map[int]bool{}
    for _, k := range nums {
        removeMap[k] = true
    }
    dummy := &ListNode{Next: head}
    cur := dummy
    for cur != nil && cur.Next != nil {
        if _, ok := removeMap[cur.Next.Val]; ok {
            cur.Next = cur.Next.Next
        } else {
            cur = cur.Next
        }
    }
    return dummy.Next
}

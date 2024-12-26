// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    // 创建一个哑节点(dummy node)，Next指向头节点，用于处理边界情况
    dummy := &ListNode{Next: head}
    // 当前指针初始化为哑节点
    cur := dummy
    
    // 遍历链表，确保当前节点及其后两个节点均不为空
    for cur != nil && cur.Next != nil && cur.Next.Next != nil {
        // 获取当前节点的下一个节点的值
        val := cur.Next.Val
        // 如果下两个节点的值相等，说明有重复
        if cur.Next.Next.Val == val {
            // 遍历并移除所有值相等的节点
            for cur.Next != nil && cur.Next.Val == val {
                cur.Next = cur.Next.Next // 删除当前重复节点
            }
        } else {
            // 否则，移动到下一个节点
            cur = cur.Next
        }
    }
    // 返回去重后的链表头
    return dummy.Next
}

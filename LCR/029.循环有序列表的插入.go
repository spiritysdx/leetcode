/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func insert(head *ListNode, insertVal int) *ListNode {
    newNode := &ListNode{Val: insertVal} // 简化节点初始化

    // 处理空链表的情况
    if head == nil {
        newNode.Next = newNode
        return newNode
    }

    // 处理只有一个节点的情况
    if head.Next == head {
        head.Next = newNode
        newNode.Next = head
        return head
    }

    cur := head
    next := head.Next

    for {
        // 找到插入点：
        // 1. 当前节点 <= 插入值 <= 下一个节点
        if insertVal >= cur.Val && insertVal <= next.Val {
            break
        }
        // 2. 处理“拐点”情况：最大值和最小值的连接处
        if cur.Val > next.Val {
            if insertVal > cur.Val || insertVal < next.Val {
                break
            }
        }

        // 更新指针
        cur = cur.Next
        next = next.Next

        // 如果到达了链表的起始点，结束循环
        if cur == head {
            break
        }
    }

    // 插入新节点
    newNode.Next = next
    cur.Next = newNode
    return head
}

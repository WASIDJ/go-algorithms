/*
 * @lc app=leetcode.cn id=86 lang=golang
 * @lcpr version=30403
 *
 * [86] 分隔链表
 */

package leetcode_solutions

import "testing"


// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    // 创建两个虚拟节点
    small := &ListNode{Val: 0}
    large := &ListNode{Val: 0}
    
    smallPtr := small
    largePtr := large
    
    // 遍历原链表
    for head != nil {
        if head.Val < x {
            smallPtr.Next = head
            smallPtr = smallPtr.Next
        } else {
            largePtr.Next = head
            largePtr = largePtr.Next
        }
        head = head.Next
    }
    
    // 连接两个链表
    largePtr.Next = nil // 防止环形链表
    smallPtr.Next = large.Next
    
    return small.Next
}
// @lc code=end

func TestPartitionList(t *testing.T) {
	// your test code here
	
}



/*
// @lcpr case=start
// [1,4,3,2,5,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n2\n
// @lcpr case=end

 */


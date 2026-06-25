/*
 * @lc app=leetcode.cn id=61 lang=golang
 * @lcpr version=30403
 *
 * [61] 旋转链表
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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// 找到链表的长度和尾节点
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// 规范化 k
	k = k % length
	if k == 0 {
		return head
	}

	// 找到新的头节点位置
	newTailPos := length - k
	newTail := head
	for i := 1; i < newTailPos; i++ {
		newTail = newTail.Next
	}

	// 旋转
	newHead := newTail.Next
	newTail.Next = nil
	tail.Next = head

	return newHead
}

// @lc code=end

func TestRotateList(t *testing.T) {
	// your test code here

}

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2]\n4\n
// @lcpr case=end

*/

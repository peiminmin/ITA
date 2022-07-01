package _go

/*
06. 从尾到头打印链表
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。



示例 1：

输入：head = [1,3,2]
输出：[2,3,1]

解题思路：
solution 1： 实现一个栈，遍历连表，将数据入栈，然后在弹出，空间复杂度O(n),时间复杂度O(n)
solution 2: 递归法
*/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	array := makeArray()
	innerReverse(head, &array)
	return array
}

func innerReverse(head *ListNode, array *[]int) {
	if head == nil {
		return
	}
	if head != nil && head.Next == nil {
		*array = append(*array, head.Val)
		return
	}

	id1, id2 := head, head.Next

	for id2.Next != nil {
		id1 = id1.Next
		id2 = id2.Next
	}
	*array = append(*array, id2.Val)
	id1.Next = nil
	innerReverse(head, array)
}

func makeArray() []int {
	return make([]int, 0)
}

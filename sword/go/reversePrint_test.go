package _go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReversePrint(t *testing.T) {
	var case1 *ListNode
	assert.Equal(t, []int(nil), reversePrint(case1), "nil case")

	case2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	assert.Equal(t, []int{2, 3, 1}, reversePrint(case2), "case2")

	case3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	assert.Equal(t, []int{4}, reversePrint(case3), "case3")
}

package _go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var case1PreOrder, caseInOrder []int
	assert.Equal(t, nil, buildTree(case1PreOrder, caseInOrder), "case1")
	case2PreOrder := []int{3, 9, 20, 15, 7}
	case2InOrder := []int{9, 3, 15, 20, 7}
	assert.Equal(t, []int{3, 9, 20, 0, 0, 15, 7}, buildTree(case2PreOrder, case2InOrder), "case2")
	case3PreOrder := []int{-1}
	case3InOrder := []int{-1}
	assert.Equal(t, []int{-1}, buildTree(case3PreOrder, case3InOrder), "case3")
}

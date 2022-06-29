package _go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	assert.Equal(t, true, findNumberIn2DArray(matrix, 5), "target 5 in matrix")
	assert.Equal(t, false, findNumberIn2DArray(matrix, 20), "target 20 not in matrix")
	matrix = [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	assert.Equal(t, true, findNumberIn2DArray(matrix, 15), "target 15 in matrix")

}

func TestFindNumber(t *testing.T) {
	test1 := []int{1, 4, 7, 11, 15}
	assert.Equal(t, false, findNumber(test1, 5), "target 5 not in matrix")
	assert.Equal(t, true, findNumber(test1, 11), "target 11  in matrix")
	test2 := make([]int, 1)
	assert.Equal(t, false, findNumber(test2, 5), "target 5 not in matrix")

}

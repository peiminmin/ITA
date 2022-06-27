package _go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestfindNumberIn2DArray(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	assert.Equal(t, true, findNumberIn2DArray(matrix, 5), "target 5 in matrix")
	assert.Equal(t, true, findNumberIn2DArray(matrix, 20), "target 20 not in matrix")
}

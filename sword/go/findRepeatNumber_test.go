package _go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindRepeatNumber(t *testing.T) {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	result := findRepeatNumber(nums)
	assert.Equal(t, 2, result, fmt.Sprintf("%d not equal expecte 2", result))
}

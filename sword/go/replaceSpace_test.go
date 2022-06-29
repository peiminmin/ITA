package _go

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	testCase1 := "We are happy."
	assert.Equal(t, "We%20are%20happy.", replaceSpace(testCase1), "case 1")
}

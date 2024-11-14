package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumberWhenMissing(t *testing.T) {
	nums := []int{3, 0, 1}
	expected := 2
	assert.Equal(t, expected, missingNumber(nums))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicatesWhenEmpty(t *testing.T) {
	nums := make([]int, 0)
	expected := false
	assert.Equal(t, expected, containsDuplicates(nums))
}

func TestContainsDuplicatesWhenNoDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	expected := false
	assert.Equal(t, expected, containsDuplicates(nums))
}

func TestContainsDuplicatesWhenDuplicates(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 4, 5, 6, 7}
	expected := true
	assert.Equal(t, expected, containsDuplicates(nums))
}

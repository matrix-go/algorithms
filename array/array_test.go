package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1. 返回数组中所有元素的总和
func calculateArraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
func TestCalculateArraySum(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}
	sum := calculateArraySum(arr[:])
	assert.Equal(t, 15, sum)
}

package stack

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 1. 一个模拟栈数据结构的实现FrequencyStack
// push(x int) 将 x 压入栈, pop 将栈中出现频次最高的元素弹出并返回
// 如果出现频次最高的元素存在多个,则返回最接近栈顶的元素
type FrequencyStack struct {
	// 元素 -> 频次
	freq map[int]int
	// 数据
	data []int
	// 当前最大的频次
	maxFreq int
}

var (
	ErrStackEmpty = errors.New("stack is empty")
)

func NewFrequencyStack() *FrequencyStack {
	return &FrequencyStack{
		freq:    make(map[int]int),
		data:    make([]int, 0),
		maxFreq: 0,
	}
}

func (s *FrequencyStack) Push(elem int) {
	s.freq[elem]++
	freq := s.freq[elem]
	s.data = append(s.data, elem)
	if freq > s.maxFreq {
		s.maxFreq = freq
	} else {
		// 交换顶元素
		if freq > 1 {
			s.data[len(s.data)-1], s.data[len(s.data)-2] = s.data[len(s.data)-2], s.data[len(s.data)-1]
		} else {
			// 从左边起第一个 freq > 1 的元素下标, 并插入对应位置
			for i := 0; i < len(s.data)-1; i++ {
				if s.freq[s.data[i]] > 1 {
					data := make([]int, len(s.data))
					copy(data, s.data[:i])
					data[i] = elem
					copy(data[i+1:], s.data[i:len(s.data)-1])
					s.data = data
					break
				}
			}
		}
	}
}

func (s *FrequencyStack) Pop() (elem int, err error) {
	if s.maxFreq == 0 {
		return 0, ErrStackEmpty
	}

	elem = s.data[len(s.data)-1]
	s.freq[elem]--
	s.data = s.data[:len(s.data)-1]

	maxFreq := 0
	for _, freq := range s.freq {
		if freq >= maxFreq {
			maxFreq = freq
		}
	}
	s.maxFreq = maxFreq

	return elem, nil
}

func TestFrequencyStack(t *testing.T) {
	stack := NewFrequencyStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(2)
	stack.Push(3)
	stack.Push(3)
	stack.Push(4)
	e1, _ := stack.Pop()
	assert.Equal(t, 3, e1)
	e2, _ := stack.Pop()
	assert.Equal(t, 2, e2)
	e3, _ := stack.Pop()
	assert.Equal(t, 3, e3)
	e4, _ := stack.Pop()
	assert.Equal(t, 3, e4)
	e5, _ := stack.Pop()
	assert.Equal(t, 2, e5)
	e6, _ := stack.Pop()
	assert.Equal(t, 4, e6)
	e7, _ := stack.Pop()
	assert.Equal(t, 1, e7)
	_, err := stack.Pop()
	assert.ErrorIs(t, err, ErrStackEmpty)
}

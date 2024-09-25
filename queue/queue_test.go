package queue

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// 1. 一个随机的人在队列中,每个人由身高和前面大于等于该身高的人的个数来表示

type PersonQueue struct {
	persons []Person
}

func NewPersonQueue() *PersonQueue {
	return &PersonQueue{persons: make([]Person, 0, 10)}
}

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

func (q *PersonQueue) Enqueue(p Person) {
	if len(q.persons) == 0 {
		q.persons = append(q.persons, p)
		return
	}
	for i := range q.persons {
		if q.persons[i].height < p.height {
			q.persons = append(q.persons, Person{})
			copy(q.persons[i+1:], q.persons[i:])
			p.count = i
			q.persons[i] = p

			for idx := range q.persons[i+1:] {
				q.persons[i+1:][idx].count++
			}
			return
		}
	}
}

func (q *PersonQueue) Dequeue() (Person, error) {
	if len(q.persons) == 0 {
		return Person{}, ErrEmptyQueue
	}
	p := q.persons[0]
	q.persons = q.persons[1:]
	for idx := range q.persons {
		q.persons[idx].count--
	}
	return p, nil
}

type Person struct {
	height int
	count  int
}

func TestPersonQueue(t *testing.T) {
	q := NewPersonQueue()
	q.Enqueue(Person{height: 1})
	assert.Equal(t, 0, q.persons[0].count)
	q.Enqueue(Person{height: 2})
	assert.Equal(t, 1, q.persons[1].count)
	q.Enqueue(Person{height: 4})
	assert.Equal(t, 2, q.persons[2].count)
	assert.Equal(t, 4, q.persons[0].height)
	q.Enqueue(Person{height: 3})
	assert.Equal(t, 3, q.persons[3].count)
	q.Enqueue(Person{height: 5})
	assert.Equal(t, 5, q.persons[0].height)
	p1, err := q.Dequeue()
	require.NoError(t, err)
	assert.Equal(t, 5, p1.height)
	assert.Equal(t, 0, p1.count)
	assert.Equal(t, 1, q.persons[1].count)
	assert.Equal(t, 2, q.persons[2].count)
	assert.Equal(t, 2, q.persons[2].height)
	p2, err := q.Dequeue()
	require.NoError(t, err)
	assert.Equal(t, 4, p2.height)
	p3, err := q.Dequeue()
	require.NoError(t, err)
	assert.Equal(t, 3, p3.height)
	p4, err := q.Dequeue()
	require.NoError(t, err)
	assert.Equal(t, 2, p4.height)
	p5, err := q.Dequeue()
	require.NoError(t, err)
	assert.Equal(t, 1, p5.height)
	_, err = q.Dequeue()
	assert.Equal(t, ErrEmptyQueue, err)
}

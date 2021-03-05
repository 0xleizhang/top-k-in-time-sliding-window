package pkg

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(5)
	assert.Equal(t, true, q.IsEmpty())
	q.EnQueue(Synopsis{})
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	assert.Equal(t, q.IsFull(), true)
	assert.Equal(t, q.Front(), 1)
	assert.Equal(t, q.Rear(), 5)
	k := q.DeQueue()
	assert.Equal(t, 1, k)
	assert.Equal(t, false, q.IsFull())
	q.EnQueue(6)
	assert.Equal(t,2,q.Front())
	q.EnQueue(12)
	b := q.EnQueue(7)
	assert.Equal(t, false, b)

}

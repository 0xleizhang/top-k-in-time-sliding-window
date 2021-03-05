package pkg

//循环队列
type CircularQueue struct {
	queue     []interface{}
	headIndex int
	count     int
	capacity  int
}

func NewCircularQueue(size int) *CircularQueue {
	q := &CircularQueue{
		queue:     make([]interface{}, size, size),
		headIndex: 0,
		count:     0,
		capacity:  size,
	}
	return q
}

/*
入队
*/
func (q *CircularQueue) EnQueue(value interface{}) bool {
	if q.count == q.capacity {
		return false
	}
	q.queue[(q.headIndex+q.count)%q.capacity] = value
	q.count += 1
	return true
}

/**
出队
*/
func (q *CircularQueue) DeQueue() interface{} {
	if q.count == 0 {
		return nil
	}
	head := q.queue[q.headIndex]
	q.headIndex = (q.headIndex + 1) % q.capacity
	q.count = q.count - 1
	return head
}

/**
队首
*/
func (q *CircularQueue) Front() interface{} {
	if q.count == 0 {
		return nil
	}
	return q.queue[q.headIndex]
}

/**
队尾
*/
func (q *CircularQueue) Rear() interface{} {
	if q.count == 0 {
		return nil
	}
	tailIndex := (q.headIndex + q.count - 1) % q.capacity
	return q.queue[tailIndex]
}

func (q *CircularQueue) IsEmpty() bool {
	return q.count == 0
}

func (q *CircularQueue) IsFull() bool {
	return q.count == q.capacity
}

/**
清空
*/
func (q *CircularQueue) Clear() {
	q.headIndex = 0
	q.count = 0
}

func (q *CircularQueue) Size() int {
	return q.count
}

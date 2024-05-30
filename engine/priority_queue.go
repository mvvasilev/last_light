package engine

import "slices"

// [ a = 1, b = 2, c = 3, d = 4 ]
// a = 1 <- [ b = 2, c = 3, d = 4 ]
// do a action
// sub a priority from queue: [ b = 1, c = 2, d = 3 ]
// [ b = 1, a = 1, c = 2, d = 3 ] <- a = 1
// b = 1 <- [ a = 1, c = 2, d = 3 ]

type priorityQueueItem[T interface{}] struct {
	priority int
	value    T
}

type PriorityQueue[T interface{}] struct {
	queue []*priorityQueueItem[T]
}

func CreatePriorityQueue[T interface{}]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		queue: make([]*priorityQueueItem[T], 0, 10),
	}
}

func (pq *PriorityQueue[T]) Enqueue(prio int, value T) {
	pq.queue = append(pq.queue, &priorityQueueItem[T]{priority: prio, value: value})
	slices.SortFunc(pq.queue, func(e1, e2 *priorityQueueItem[T]) int { return e1.priority - e2.priority })
}

func (pq *PriorityQueue[T]) AdjustPriorities(amount int) {
	for _, e := range pq.queue {
		e.priority += amount
	}
}

// Pop element w/ lowest priority
func (pq *PriorityQueue[T]) DequeueValue() (value T) {
	if len(pq.queue) < 1 {
		return
	}

	value, pq.queue = pq.queue[0].value, pq.queue[1:len(pq.queue)]

	return
}

// Peek element w/ lowest priority
func (pq *PriorityQueue[T]) Peek() (priority int, value T) {
	if len(pq.queue) < 1 {
		return
	}

	priority, value = pq.queue[0].priority, pq.queue[0].value

	return
}

// Pop element w/ lowest priority, returning the priority as well as the value
func (pq *PriorityQueue[T]) Dequeue() (priority int, value T) {
	if len(pq.queue) < 1 {
		return
	}

	priority, value, pq.queue = pq.queue[0].priority, pq.queue[0].value, pq.queue[1:len(pq.queue)]

	return
}

func (pq *PriorityQueue[T]) Clear() {
	pq.queue = make([]*priorityQueueItem[T], 0, 10)
}

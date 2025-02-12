package queue

// Queue represents a generic FIFO queue data structure.
// Elements are added to the back and removed from the front.
// The zero value is not usable; use NewQueue to create a new Queue.
type Queue[T any] struct {
	elements []T
}

// NewQueue creates and returns an empty queue that can store elements of type T.
//
// Example:
//
//	q := NewQueue[int]()
//	q.Enqueue(1)
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elements: make([]T, 0),
	}
}

// Enqueue adds an element to the back of the queue.
//
// Example:
//
//	q := NewQueue[int]()
//	q.Enqueue(1) // queue now contains: [1]
//	q.Enqueue(2) // queue now contains: [1, 2]
func (q *Queue[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

// Dequeue removes and returns the element at the front of the queue.
// Returns the element and true if successful, or zero value and false if the queue is empty.
//
// Example:
//
//	q := NewQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	val, ok := q.Dequeue() // val = 1, ok = true
//	val, ok = q.Dequeue()  // val = 2, ok = true
//	val, ok = q.Dequeue()  // val = 0, ok = false (queue empty)
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var empty T
		return empty, false
	}

	element := q.elements[0]

	if q.Length() == 1 {
		// Only one element remaining. Reset the queue to prevent memory leaks
		q.elements = nil

		return element, true
	}

	// remove element from queue
	q.elements = q.elements[1:]

	return element, true
}

// Length returns the number of elements currently in the queue.
//
// Example:
//
//	q := NewQueue[int]()
//	q.Enqueue(1)
//	q.Enqueue(2)
//	fmt.Println(q.Length()) // Output: 2
func (q *Queue[T]) Length() int {
	return len(q.elements)
}

// IsEmpty returns true if the queue contains no elements, false otherwise.
//
// Example:
//
//	q := NewQueue[int]()
//	fmt.Println(q.IsEmpty()) // Output: true
//	q.Enqueue(1)
//	fmt.Println(q.IsEmpty()) // Output: false
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

// Peek returns the element at the front of the queue without removing it.
// Returns the element and true if successful, or zero value and false if the queue is empty.
//
// Example:
//
//	q := NewQueue[int]()
//	q.Enqueue(1)
//	val, ok := q.Peek() // val = 1, ok = true, queue still contains: [1]
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var empty T
		return empty, false
	}

	return q.elements[0], true
}
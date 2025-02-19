package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	var v int
	var ok bool

	queue := NewQueue[int]()
	assertEquals(t, queue.Length(), 0)
	assertEquals(t, queue.IsEmpty(), true)

	v, ok = queue.Peek()
	assertEquals(t, ok, false)

	queue.Enqueue(10)
	assertEquals(t, queue.Length(), 1)
	assertEquals(t, queue.IsEmpty(), false)

	v, ok = queue.Peek()
	assertEquals(t, queue.Length(), 1)
	assertEquals(t, v, 10)
	assertEquals(t, ok, true)

	queue.Enqueue(20)
	assertEquals(t, queue.Length(), 2)
	assertEquals(t, queue.IsEmpty(), false)

	v, ok = queue.Peek()
	assertEquals(t, queue.Length(), 2)
	assertEquals(t, v, 10)
	assertEquals(t, ok, true)

	v, ok = queue.Dequeue()
	assertEquals(t, queue.Length(), 1)
	assertEquals(t, queue.IsEmpty(), false)
	assertEquals(t, ok, true)
	assertEquals(t, v, 10)

	v, ok = queue.Peek()
	assertEquals(t, queue.Length(), 1)
	assertEquals(t, v, 20)
	assertEquals(t, ok, true)

	v, ok = queue.Dequeue()
	assertEquals(t, queue.Length(), 0)
	assertEquals(t, queue.IsEmpty(), true)
	assertEquals(t, ok, true)
	assertEquals(t, v, 20)

	v, ok = queue.Dequeue()
	assertEquals(t, queue.Length(), 0)
	assertEquals(t, queue.IsEmpty(), true)
	assertEquals(t, ok, false)

	queue.Enqueue(30)
	assertEquals(t, queue.Length(), 1)
	assertEquals(t, queue.IsEmpty(), false)
}

func TestQueue_PreventDuplicates(t *testing.T) {
	type ContactUser struct {
		Email string
	}

	queue := NewQueue[ContactUser]()
	err := queue.PreventDuplicates(func(a, b ContactUser) bool {
		return a.Email == b.Email
	})
	assertEquals(t, queue.Length(), 0)
	assertEquals(t, queue.IsEmpty(), true)
	assertEquals(t, err, nil)

	queue.Enqueue(ContactUser{Email: "alice@example.com"})
	assertEquals(t, queue.Length(), 1)

	queue.Enqueue(ContactUser{Email: "bob@example.com"})
	assertEquals(t, queue.Length(), 2)

	queue.Enqueue(ContactUser{Email: "alice@example.com"})
	assertEquals(t, queue.Length(), 2)

	queueNotComparable := NewQueue[any]()
	err = queueNotComparable.PreventDuplicates(func(a, b any) bool {
		return false
	})
	if err == nil {
		t.Errorf("failed to return error")
	}
}

func assertEquals[V comparable](t *testing.T, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

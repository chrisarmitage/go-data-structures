package set

import "sync"

// Set represents a thread-safe collection of unique elements.
// The zero value is not usable; use NewSet to create a new Set.
type Set[T comparable] struct {
	members map[T]struct{}
	mu      sync.RWMutex
}

// NewSet creates and initializes a new empty Set.
//
// Example:
//
//	s := NewSet[string]()
//	s.Add("foo")
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		members: make(map[T]struct{}),
	}
}

// Members returns a slice containing all elements in the Set.
// The order of elements is not guaranteed to be stable between calls.
//
// Example:
//
//	s := NewSet[int]()
//	s.Add(1)
//	s.Add(2)
//	fmt.Println(s.Members()) // Output: [1 2] (order not guaranteed)
func (s *Set[T]) Members() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	members := make([]T, 0, len(s.members))
	for member := range s.members {
		members = append(members, member)
	}
	return members
}

// Add inserts an element into the Set.
// If the element already exists, the Set remains unchanged.
//
// Example:
//
//	s := NewSet[int]()
//	s.Add(1) // Set now contains 1
//	s.Add(1) // Set still contains just 1
func (s *Set[T]) Add(member T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.members[member] = struct{}{}
}

// Remove deletes an element from the Set.
// If the element doesn't exist, the Set remains unchanged.
// This operation is thread-safe.
//
// Example:
//
//	s := NewSet[int]()
//	s.Add(1)
//	s.Remove(1) // Set is now empty
//	s.Remove(1) // No effect - element wasn't present
func (s *Set[T]) Remove(member T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.members, member)
}

// Contains returns true if the element exists in the Set, false otherwise.
// This operation is thread-safe.
//
// Example:
//
//	s := NewSet[string]()
//	s.Add("foo")
//	fmt.Println(s.Contains("foo")) // Output: true
//	fmt.Println(s.Contains("bar")) // Output: false
func (s *Set[T]) Contains(member T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.members[member]
	return exists
}

// Size returns the number of elements in the Set.
// This operation is thread-safe.
//
// Example:
//
//	s := NewSet[int]()
//	s.Add(1)
//	s.Add(2)
//	fmt.Println(s.Size()) // Output: 2
func (s *Set[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.members)
}

// Clear removes all elements from the Set.
// This operation is thread-safe.
//
// Example:
//
//	s := NewSet[int]()
//	s.Add(1)
//	s.Add(2)
//	s.Clear() // Set is now empty
//	fmt.Println(s.Len()) // Output: 0
func (s *Set[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.members = make(map[T]struct{})
}

// Intersect returns a new set containing elements that are present in both sets.
// This operation is thread-safe and does not modify the original sets.
//
// Example:
//
//	s1 := NewSet[int]()
//	s1.Add(1)
//	s1.Add(2)
//	s2 := NewSet[int]()
//	s2.Add(2)
//	s2.Add(3)
//	result := s1.Intersect(s2)
//	fmt.Println(result.Members()) // Output: [2]
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for member := range s.members {
		if _, exists := other.members[member]; exists {
			result.Add(member)
		}
	}
	return result
}

// Union returns a new set containing all elements from both sets.
// This operation is thread-safe and does not modify the original sets.
//
// Example:
//
//	s1 := NewSet[int]()
//	s1.Add(1)
//	s1.Add(2)
//	s2 := NewSet[int]()
//	s2.Add(2)
//	s2.Add(3)
//	result := s1.Union(s2)
//	fmt.Println(result.Members()) // Output: [1 2 3]
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for member := range s.members {
		result.Add(member)
	}
	for member := range other.members {
		result.Add(member)
	}
	return result
}

// Difference returns a new set containing elements that are present in the current set but not in the other set.
// This operation is thread-safe and does not modify the original sets.
//
// Example:
//
//	s1 := NewSet[int]()
//	s1.Add(1)
//	s1.Add(2)
//	s1.Add(3)
//	s2 := NewSet[int]()
//	s2.Add(2)
//	s2.Add(3)
//	s2.Add(4)
//	result := s1.Difference(s2)
//	fmt.Println(result.Members()) // Output: [1]
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	s.mu.RLock()
	defer s.mu.RUnlock()
	other.mu.RLock()
	defer other.mu.RUnlock()
	for member := range s.members {
		if _, exists := other.members[member]; !exists {
			result.Add(member)
		}
	}
	return result
}

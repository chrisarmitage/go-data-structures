package set

import (
	"slices"
	"testing"
)

func TestSet_AddRemoveSize(t *testing.T) {
	set := NewSet[int]()
	assertEquals(t, set.Size(), 0)

	set.Add(1)
	assertEquals(t, set.Size(), 1)

	set.Add(2)
	assertEquals(t, set.Size(), 2)

	set.Add(1)
	assertEquals(t, set.Size(), 2)
	assertEquals(t, set.Contains(1), true)

	set.Remove(1)
	assertEquals(t, set.Size(), 1)
	assertEquals(t, set.Contains(1), false)
}

func TestSet_Members(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Add(3)

	members := set.Members()
	assertEquals(t, len(members), 3)
	assertEquals(t, slices.Contains(members, 1), true)
	assertEquals(t, slices.Contains(members, 2), true)
	assertEquals(t, slices.Contains(members, 3), true)
}

func TestSet_Clear(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	assertEquals(t, set.Size(), 3)

	set.Clear()

	assertEquals(t, set.Size(), 0)
	assertEquals(t, set.Contains(1), false)
	assertEquals(t, set.Contains(2), false)
	assertEquals(t, set.Contains(3), false)
}

func TestIntersect(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewSet[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	result := s1.Intersect(s2)
	members := result.Members()

	assertEquals(t, len(members), 2)
	assertEquals(t, slices.Contains(members, 1), false)
	assertEquals(t, slices.Contains(members, 2), true)
	assertEquals(t, slices.Contains(members, 3), true)
	assertEquals(t, slices.Contains(members, 4), false)
}

func TestUnion(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewSet[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	result := s1.Union(s2)
	members := result.Members()

	assertEquals(t, len(members), 4)
	assertEquals(t, slices.Contains(members, 1), true)
	assertEquals(t, slices.Contains(members, 2), true)
	assertEquals(t, slices.Contains(members, 3), true)
	assertEquals(t, slices.Contains(members, 4), true)
}

func TestDifference(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewSet[int]()
	s2.Add(2)
	s2.Add(3)
	s2.Add(4)

	result := s1.Difference(s2)
	members := result.Members()

	assertEquals(t, len(members), 1)
	assertEquals(t, slices.Contains(members, 1), true)
	assertEquals(t, slices.Contains(members, 2), false)
	assertEquals(t, slices.Contains(members, 3), false)
	assertEquals(t, slices.Contains(members, 4), false)
}

func assertEquals[V comparable](t *testing.T, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

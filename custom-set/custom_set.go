package stringset

import (
	"fmt"
	"strings"
)

/*TestVersion is the unit test this will pass.*/
const TestVersion = 2

/*Set is an unordered collection of items with only a single copy of each item.*/
type Set map[string]bool

/*New creats a new empty set*/
func New() Set {
	return Set{}
}

/*NewFromSlice creates a new set containing the elements in a slice.*/
func NewFromSlice(slice []string) Set {
	s := New()
	for _, v := range slice {
		s.Add(v)
	}
	return s
}

/*Add adds an item to the set.*/
func (s Set) Add(str string) {
	if !s.Has(str) {
		s[str] = true
	}
}

/*Delete removes an item from the set.*/
func (s Set) Delete(str string) {
	delete(s, str)
}

/*Has is true if the item is in the set.*/
func (s Set) Has(str string) bool {
	i, ok := s[str]
	return ok && i
}

/*IsEmpty is true if there are no items in the set.*/
func (s Set) IsEmpty() bool {
	return s.Len() == 0
}

/*Len counts the numer of items in the set.*/
func (s Set) Len() int {
	return len(s)
}

/*Slice formats a set as an array slice.*/
func (s Set) Slice() []string {
	keys := []string{}
	for key := range s {
		keys = append(keys, key)
	}
	return keys
}

/*String formats a set as a string.*/
func (s Set) String() string {
	var items []string
	for _, v := range s.Slice() {
		items = append(items, fmt.Sprintf("\"%v\"", v))
	}
	return "{" + strings.Join(items, ", ") + "}"
}

/*Equal determines if two sets have all elements in common.*/
func Equal(s1, s2 Set) bool {
	return SymmetricDifference(s1, s2).Len() == 0
}

/*Subset determines if the first set has all elemets in common with the second.*/
func Subset(s1, s2 Set) bool {
	return Equal(s1, Intersection(s1, s2))
}

/*Disjoint determines if two sets have no elements in common.*/
func Disjoint(s1, s2 Set) bool {
	return Intersection(s1, s2).Len() == 0
}

/*Intersection caluclates the set of elements two set have in common.*/
func Intersection(s1, s2 Set) Set {
	intersection := New()
	for _, v := range s1.Slice() {
		if s2.Has(v) {
			intersection.Add(v)
		}
	}
	return intersection
}

/*Union calculates the set of elements that either of two sets have in common.*/
func Union(s1, s2 Set) Set {
	union := New()
	for _, v := range append(s1.Slice(), s2.Slice()...) {
		union.Add(v)
	}
	return union
}

/*Difference calculates the set of elements in the first set but not in the second.*/
func Difference(s1, s2 Set) Set {
	difference := New()
	for _, v := range s1.Slice() {
		if !s2.Has(v) {
			difference.Add(v)
		}
	}
	return difference
}

/*SymmetricDifference calculates the set of elements that are only on one set.*/
func SymmetricDifference(s1, s2 Set) Set {
	difference := Union(s1, s2)
	for _, v := range Intersection(s1, s2).Slice() {
		difference.Delete(v)
	}
	return difference
}

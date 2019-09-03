package lib

// IntSet is a set of Ints
type IntSet struct {
  Set map[int]bool
}

// NewIntSet creates a new initialized IntSet
func NewIntSet() IntSet {
  return IntSet{Set: map[int]bool{}}
}

// Add adds an element to a set
func (set *IntSet) Add(i int) bool {
  _, found := set.Set[i]
  set.Set[i] = true
  return !found
}

// StringSet is a set of Strings
type StringSet struct {
  Set map[string]bool
}

// NewStringSet creates a new initialized StringSet
func NewStringSet() StringSet {
  return StringSet{Set: map[string]bool{}}
}

// Add an element to a set
func (set *StringSet) Add(s string) bool {
  _, found := set.Set[s]
  set.Set[s] = true
  return !found
}

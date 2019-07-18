package intset

import (
	"bytes"
	"fmt"
)

// IntSet defines an integer set for small positive integers
type IntSet struct {
	words []uint
}

// Wordsize is the word size of the architecture 32 or 64 bit
const (
	WordSize = 32 << (^uint(0) >> 63)
)

// Has reports whether a set contains non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/WordSize, x%WordSize
	return word < len(s.words) && s.words[word]&(1<<uint(bit)) != 0
}

// Add adds a non-negative integer to the set
func (s *IntSet) Add(x int) {
	word, bit := x/WordSize, x%WordSize
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << uint(bit)
}

// UnionWith sets s to the ubion of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", WordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements ex 6.1
func (s *IntSet) Len() (len int) {
	for _, word := range s.words {
		len += popCount(word)
	}
	return len
}

// Remove x from the set of elements ex 6.1
func (s *IntSet) Remove(x int) {
	word, bit := x/WordSize, x%WordSize
	if s.Has(x) != true {
		return
	}
	s.words[word] ^= 1 << uint(bit)
}

// Clear removes all elements from the set ex 6.1
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy returns a copy of the set ex 6.1
func (s *IntSet) Copy() *IntSet {
	var c IntSet
	c.words = make([]uint, len(s.words))
	for i, word := range s.words {
		c.words[i] = word
	}
	return &c
}

// AddAll method allwos a list of values to be added, such as s.AddAll(1,2,3)
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// IntersectionWith caclulate the bit wise intersection of sets.
func (s *IntSet) IntersectionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith calulates the bit wise difference of two set.
// i.e. delete all elemets of set1 where set 2 is set. This is not
// to be confused with the ^ operator
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		}
	}
}

// SymmetricDifference calculates the differnce of to sets where only the
// elements are present in one or the other set, the exclusive or operator ^.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems returns the integer elements of the set as a slize
func (s *IntSet) Elems() []int {
	setsize := s.Len()
	slice := make([]int, setsize)
	var k int

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < WordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				slice[k] = i*WordSize + j
				k++
			}
		}
	}
	return slice
}

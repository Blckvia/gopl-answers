// ex6.1 create Len, Remove, Clear, Copy methods
package main

type IntSet struct {
	words []uint64
	count int
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Len() int {
	return s.count
}

func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/64, uint(x%64)
		s.words[word] &^= 1 << bit
		s.count--
	}
}

func (s *IntSet) Clear() {
	s.words = nil
	s.count = 0
}

func (s *IntSet) Copy() *IntSet {
	new := &IntSet{}
	new.words = make([]uint64, len(s.words))
	copy(new.words, s.words)
	return new
}

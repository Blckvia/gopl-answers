package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type HashSet map[int]bool

func newHashSet() HashSet {
	return make(map[int]bool)
}

func (s HashSet) Has(x int) bool {
	return s[x]
}

func (s HashSet) Add(x int) {
	s[x] = true
}

func (s HashSet) UnionWith(t HashSet) {
	for v := range t {
		s[v] = true
	}
}

func (s HashSet) String() string {
	var buf bytes.Buffer
	var items []int
	for v := range s {
		items = append(items, v)
	}

	sort.Ints(items)
	buf.WriteByte('{')
	for i, item := range items {
		fmt.Fprintf(&buf, "%d", item)
		if i != len(items)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

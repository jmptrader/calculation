// Package sym provides symbol table functionality.
package calc

type Tbl struct {
	words map[string]int
	syms  map[int]string
}

// Add adds a word to the table.
func (t *Tbl) Add(word string) {
	_, ok := t.words[word]
	if !ok {
		i := len(t.words) + 1
		t.words[word] = i
		t.syms[i] = word
	}
}

package trie

//https://en.wikipedia.org/wiki/Trie

type Node struct {
	Char     rune
	Children map[rune]*Node
}

func NewNode(char rune) *Node {
	node := &Node{Char: char}
	node.Children = make(map[rune]*Node)
	return node
}

type Trie struct {
	RootNode *Node
}

func NewTrie() *Trie {
	// this is not used!
	root := NewNode(0)
	return &Trie{RootNode: root}
}

func (t *Trie) Insert(word string) error {
	current := t.RootNode
	for _, char := range word {
		if _, ok := current.Children[char]; !ok {
			current.Children[char] = NewNode(char)
		}
		current = current.Children[char]

	}
	// set word and final
	return nil
}

func (t *Trie) SearchWord(word string) bool {
	current := t.RootNode
	for _, char := range word {
		if current == nil {
			return false
		} else if _, ok := current.Children[char]; !ok {
			return false
		}
	}
	return true
}

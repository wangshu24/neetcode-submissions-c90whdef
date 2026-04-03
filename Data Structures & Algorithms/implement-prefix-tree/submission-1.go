type PrefixTree struct {
	root *TrieNode
}

type TrieNode struct {
	children [26] *TrieNode
	end bool
}

func Constructor() PrefixTree {
    return PrefixTree{}
}

func (this *PrefixTree) Insert(word string) {
	if this.root == nil {
		newRoot := TrieNode {
			children : [26]*TrieNode{},
			end: false,
		}
		this.root = &newRoot
	}
	tmp := this.root
	for _, val := range word {
		if tmp.children[val-'a'] == nil {
			tmp.children[val-'a'] = &TrieNode{
				children : [26]*TrieNode{},
				end: false,
			}
		} 
		tmp = tmp.children[val-'a']
	}
	tmp.end = true
}

func (this *PrefixTree) Search(word string) bool {
	if this.root == nil {return false}
	temp := this.root
	for _,val := range word {
		if temp.children[val-'a'] != nil {
			temp = temp.children[val-'a']
		} else {return false}
	} 
	if temp.end {return true}
	return false
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	if this.root == nil {return false}
	temp := this.root
	for _,val := range prefix {
		if temp.children[val-'a'] != nil {
			temp = temp.children[val-'a']
		} else {return false}
	} 
	return true
}

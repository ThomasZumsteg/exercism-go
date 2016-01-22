package binarysearchtree

//SearchTreeData is a binary tree data structure
type SearchTreeData struct {
	data        int
	left, right *SearchTreeData
}

/*Bst creates a new binary search tree.*/
func Bst(data int) *SearchTreeData {
	return &SearchTreeData{data: data}
}

/*Insert adds a node to the binary tree.*/
func (t *SearchTreeData) Insert(data int) {
	node := &t.left
	if t.data < data {
		node = &t.right
	}
	if *node == nil {
		*node = Bst(data)
	} else {
		(*node).Insert(data)
	}
}

/*MapString converts the binary tree to an array
and applys a function to every element.*/
func (t *SearchTreeData) MapString(stringFunc func(int) string) []string {
	var strings []string
	for _, i := range t.MapInt(func(i int) int { return i }) {
		strings = append(strings, stringFunc(i))
	}
	return strings
}

/*MapInt converts a binary tree to an array
and applys a function to every element.*/
func (t *SearchTreeData) MapInt(intFunc func(int) int) []int {
	if t == nil {
		return []int{}
	}
	right := append([]int{intFunc(t.data)}, t.right.MapInt(intFunc)...)
	return append(t.left.MapInt(intFunc), right...)
}

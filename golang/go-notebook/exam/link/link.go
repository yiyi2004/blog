package link

// Node -
type Node struct {
	key    int
	value  interface{}
	parent *Node
	left   *Node
	right  *Node
}

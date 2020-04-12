package llrb

import "math"

type Color bool

const (
	BLACK Color = true
	RED   Color = false
)

// Key is base interface for abstraction
type Key interface{ Compare(with Key) int }

// Value is base interface for abstraction
type Value interface{}

// Node ...
type Node struct {
	left, right *Node
	key         Key
	value       Value
	color       Color
}

func isRed(n *Node) bool {
	if n == nil {
		return false
	}
	return n.color == RED
}

func newNode(k Key, v Value) *Node {
	return &Node{key: k, value: v, color: RED}
}

// Tree ...
type Tree struct {
	root *Node
	size int
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) insert(k Key, v Value) {
	t.root = insertHelper(t.root, k, v)
	t.root.color = BLACK
}

func insertHelper(n *Node, k Key, v Value) *Node {
	if n == nil {
		return newNode(k, v)
	}

	// Insertion
	cmp := n.key.Compare(k)
	if cmp == 0 {
		n.value = v
	}
	if cmp < 0 {
		n.left = insertHelper(n.left, k, v)
	}
	if cmp > 0 {
		n.right = insertHelper(n.right, k, v)
	}

	// Balancing
	if isRed(n.right) && !isRed(n.left) {
		n = rotateLeft(n)
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}
	if isRed(n.left) && isRed(n.right) {
		colorFlip(n)
	}
	return n
}

func rotateLeft(n *Node) *Node {
	x := n.right
	n.right = x.left
	x.left = n

	x.color = n.color
	n.color = RED

	return x
}

func rotateRight(n *Node) *Node {
	x := n.left
	n.left = x.right
	x.right = n

	x.color = n.color
	n.color = RED

	return x
}

func colorFlip(n *Node) {
	n.color = !n.color
	n.left.color = !n.left.color
	n.right.color = !n.right.color
}

// Search searches given key in the tree and returns nil if search fails
func (t *Tree) search(k Key) Value {
	x := t.root
	for x != nil {
		cmp := x.key.Compare(k)
		if cmp == 0 {
			return x.value
		}
		if cmp < 0 {
			x = x.left
		}
		if cmp > 0 {
			x = x.right
		}
	}
	return nil
}

func isBalanced(n *Node) bool {
	if n == nil {
		return false
	}

	lh := heightOf(n.left)
	rh := heightOf(n.right)

	if math.Abs(float64(lh-rh)) <= 1 &&
		isBalanced(n.left) &&
		isBalanced(n.right) {
		return true
	}
	return false
}

func heightOf(n *Node) int {
	if n == nil {
		return 0
	}

	lh := float64(heightOf(n.left))
	rh := float64(heightOf(n.right))
	return 1 + int(math.Max(lh, rh))
}

// nolint
package avl

import (
	"bytes"
	"fmt"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type AVLNode struct {
	Left, Right *AVLNode
	Data        int
	Height      int
}

func height(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func rotateRight(y *AVLNode) *AVLNode {
	x := y.Left
	t := x.Right
	x.Right = y
	y.Left = t
	y.Height = maxInt(height(y.Left), height(y.Right)) + 1
	x.Height = maxInt(height(x.Left), height(x.Right)) + 1
	return x
}

func rotateLeft(x *AVLNode) *AVLNode {
	y := x.Right
	t := y.Left
	y.Left = x
	x.Right = t
	x.Height = maxInt(height(x.Left), height(x.Right)) + 1
	y.Height = maxInt(height(y.Left), height(y.Right)) + 1
	return y
}

func balance(n *AVLNode) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

func Insert(root *AVLNode, data int) *AVLNode {
	if root == nil {
		return &AVLNode{Left: nil, Right: nil, Data: data, Height: 1}
	}

	if root.Data > data {
		root.Left = Insert(root.Left, data)
	} else if root.Data < data {
		root.Right = Insert(root.Right, data)
	} else {
		return root
	}

	root.Height = maxInt(height(root.Left), height(root.Right)) + 1

	balance := balance(root)
	switch {
	case balance > 1:
		if root.Left.Data > data {
			return rotateRight(root)
		} else if root.Left.Data < data {
			root.Left = rotateLeft(root.Left)
			return rotateRight(root)
		}
	case balance < -1:
		if root.Right.Data > data {
			root.Right = rotateRight(root.Right)
			return rotateLeft(root)
		} else if root.Right.Data < data {
			return rotateLeft(root)
		}
	}
	return root
}

func InOrder(root *AVLNode) string {
	var buf bytes.Buffer
	if root != nil {
		buf.WriteString(InOrder(root.Left))
		buf.WriteString(fmt.Sprintf("%d ", root.Data))
		buf.WriteString(InOrder(root.Right))
	}
	return buf.String()
}

func PreOrder(root *AVLNode) string {
	var buf bytes.Buffer
	if root != nil {
		buf.WriteString(fmt.Sprintf("%d ", root.Data))
		buf.WriteString(PreOrder(root.Left))
		buf.WriteString(PreOrder(root.Right))
	}
	return buf.String()
}

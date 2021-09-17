package avl

import (
	"strings"
	"testing"
)

func TestAVL(t *testing.T) {
	ds := []int{10, 20, 30, 40, 50, 25}
	var root *AVLNode
	for _, v := range ds {
		root = Insert(root, v)
	}

	if out := strings.TrimSpace(InOrder(root)); out != "10 20 25 30 40 50" {
		t.Error("invalid in-order:", out)
	}
	if out := strings.TrimSpace(PreOrder(root)); out != "30 20 10 25 40 50" {
		t.Error("invalid pre-order:", out)
	}
}

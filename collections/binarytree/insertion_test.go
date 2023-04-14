package binarytree

import "testing"

func TestXxx(t *testing.T) {
	var root Node[int64] = &NodeImpl[int64]{10, nil, nil, nil}
	var valueNode Node[int64] = &NodeImpl[int64]{11, nil, nil, nil}

	result := Insert(root, valueNode, CompareInt64To)

	if result.GetParentNode() == nil {
		t.Error("Parent should not be nil")
	}
	if root.GetRightNode() == nil {
		t.Error("Root right node should not be nil")
	}
	valueNode = &NodeImpl[int64]{12, nil, nil, nil}

	result = Insert(root, valueNode, CompareInt64To)
	if result.GetParentNode().GetValue() != 11 {
		t.Error("Nil not valid")
	}
}

func CompareInt64To(a, b int64) int64 {
	return a - b
}

package binarytree

import (
	"collections/common"
	"collections/tree"
)

func insertRecursive[T comparable](parent *tree.Node[T], value T, compareFunc common.CompareFunc[T]) *tree.Node[T] {
	if compareFunc(value, parent.Value) > 0 {
		if parent.Next == nil {
			parent.Next = tree.NewNode(value, parent)
			return parent.Next
		} else {
			return insertRecursive(parent.Next, value, compareFunc)
		}
	} else {
		if parent.Prev == nil {
			parent.Prev = tree.NewNode(value, parent)
			return parent.Prev
		} else {
			return insertRecursive(parent.Prev, value, compareFunc)
		}
	}
}

func Insert[T comparable](treeObj *tree.Tree[T], value T, compareFunc common.CompareFunc[T]) *tree.Node[T] {
	if treeObj.Root == nil {
		treeObj.Root = tree.NewNode(value, nil)
		return treeObj.Root
	} else {
		return insertRecursive(treeObj.Root, value, compareFunc)
	}
}

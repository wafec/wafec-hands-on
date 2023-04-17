package redblacktree

import (
	"collections/common"
	"collections/tree"
	"collections/tree/binarytree"
	"collections/tree/rotation"
)

func Insert[T comparable](treeObj *tree.Tree[T], value T, compareFunc common.CompareFunc[T], metaObj Meta[T]) *tree.Node[T] {
	nodeInserted := binarytree.Insert(treeObj, value, compareFunc)
	metaObj.SetColor(nodeInserted, Red)
	if nodeInserted == treeObj.Root {
		metaObj.SetColor(nodeInserted, Black)
	} else {
		checkUncleColor(nodeInserted, metaObj)
	}
	return nodeInserted
}

func checkUncleColor[T comparable](nodeInserted *tree.Node[T], metaObj Meta[T]) {
	uncleObj := nodeInserted.GetUncleOrNil()
	if uncleObj == nil {
		return
	}
	if metaObj.GetColor(uncleObj) == Red {
		metaObj.SetColor(uncleObj, Black)
		metaObj.SetColor(nodeInserted.Parent, Black)
		metaObj.SetColor(uncleObj.Parent, Red)
		checkUncleColor(uncleObj.Parent, metaObj)
	} else {
		doRotation(nodeInserted, metaObj)
	}
}

func doRotation[T comparable](nodeInserted *tree.Node[T], metaObj Meta[T]) {
	if isLeftLeftCase(nodeInserted) {
		doLeftLeftCase(nodeInserted, metaObj)
	} else if isLeftRightCase(nodeInserted) {
		doLeftRightCase(nodeInserted, metaObj)
	} else if isRightRightCase(nodeInserted) {
		doRightRightCase(nodeInserted, metaObj)
	} else if isRightLeftCase(nodeInserted) {
		doRightLeftCase(nodeInserted, metaObj)
	}
}

func doLeftLeftCase[T comparable](nodeInserted *tree.Node[T], metaObj Meta[T]) {
	metaObj.SetColor(nodeInserted.GetGrandFatherOrNil(), Red)
	metaObj.SetColor(nodeInserted.Parent, Black)
	rotation.RotateToRight(nodeInserted.Parent)
}

func doLeftRightCase[T comparable](nodeInserted *tree.Node[T], metaObj Meta[T]) {
	metaObj.SetColor(nodeInserted, Black)
	metaObj.SetColor(nodeInserted.GetGrandFatherOrNil(), Red)
	rotation.RotateToLeft(nodeInserted)
	rotation.RotateToRight(nodeInserted.Parent)
}

func doRightRightCase[T comparable](nodeInserted *tree.Node[T], metaObj Meta[T]) {
	metaObj.SetColor(nodeInserted.Parent, Black)
	metaObj.SetColor(nodeInserted.GetGrandFatherOrNil(), Red)
	rotation.RotateToLeft(nodeInserted.Parent)
}

func doRightLeftCase[T comparable](nodeInserted *tree.Node[T], metaObj Meta[T]) {
	metaObj.SetColor(nodeInserted, Black)
	metaObj.SetColor(nodeInserted.GetGrandFatherOrNil(), Red)
	rotation.RotateToRight(nodeInserted)
	rotation.RotateToLeft(nodeInserted.Parent)
}

func isLeftLeftCase[T comparable](nodeInserted *tree.Node[T]) bool {
	parent := nodeInserted.Parent
	grandFather := parent.Parent
	return grandFather.Prev == parent && parent.Prev == nodeInserted
}

func isLeftRightCase[T comparable](nodeInserted *tree.Node[T]) bool {
	parent := nodeInserted.Parent
	grandFather := parent.Parent
	return grandFather.Prev == parent && parent.Next == nodeInserted
}

func isRightRightCase[T comparable](nodeInserted *tree.Node[T]) bool {
	parent := nodeInserted.Parent
	grandFather := parent.Parent
	return grandFather.Next == parent && parent.Next == nodeInserted
}

func isRightLeftCase[T comparable](nodeInserted *tree.Node[T]) bool {
	parent := nodeInserted.Parent
	grandFather := parent.Parent
	return grandFather.Next == parent && parent.Prev == nodeInserted
}

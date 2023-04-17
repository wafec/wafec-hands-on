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
	checkRootColor(nodeInserted, metaObj)
	treeObj.UpdateRoot()
	return nodeInserted
}

func checkRootColor[T comparable](currentNode *tree.Node[T], metaObj Meta[T]) {
	if currentNode.IsRoot() {
		metaObj.SetColor(currentNode, Black)
	} else {
		checkUncleColor(currentNode, metaObj)
	}
}

func checkUncleColor[T comparable](currentNode *tree.Node[T], metaObj Meta[T]) {
	uncleObj := currentNode.GetUncleOrNil()
	if uncleObj != nil && metaObj.GetColor(uncleObj) == Red {
		metaObj.SetColor(uncleObj, Black)
		metaObj.SetColor(currentNode.Parent, Black)
		metaObj.SetColor(uncleObj.Parent, Red)
		checkRootColor(uncleObj.Parent, metaObj)
	} else {
		doRotation(currentNode, metaObj)
	}
}

func doRotation[T comparable](nodeInserted *tree.Node[T], metaObj Meta[T]) {
	if nodeInserted.GetGrandFatherOrNil() == nil {
		return
	}
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

func doLeftLeftCase[T comparable](currentNode *tree.Node[T], metaObj Meta[T]) {
	metaObj.SetColor(currentNode.GetGrandFatherOrNil(), Red)
	metaObj.SetColor(currentNode.Parent, Black)
	rotation.RotateToRight(currentNode.Parent)
}

func doLeftRightCase[T comparable](currentNode *tree.Node[T], metaObj Meta[T]) {
	metaObj.SetColor(currentNode, Black)
	metaObj.SetColor(currentNode.GetGrandFatherOrNil(), Red)
	rotation.RotateToLeft(currentNode)
	rotation.RotateToRight(currentNode)
}

func doRightRightCase[T comparable](currentNode *tree.Node[T], metaObj Meta[T]) {
	metaObj.SetColor(currentNode.Parent, Black)
	metaObj.SetColor(currentNode.GetGrandFatherOrNil(), Red)
	rotation.RotateToLeft(currentNode.Parent)
}

func doRightLeftCase[T comparable](currentNode *tree.Node[T], metaObj Meta[T]) {
	metaObj.SetColor(currentNode, Black)
	metaObj.SetColor(currentNode.GetGrandFatherOrNil(), Red)
	rotation.RotateToRight(currentNode)
	rotation.RotateToLeft(currentNode)
}

func isLeftLeftCase[T comparable](currentNode *tree.Node[T]) bool {
	parent := currentNode.Parent
	grandFather := parent.Parent
	return grandFather.Prev == parent && parent.Prev == currentNode
}

func isLeftRightCase[T comparable](currentNode *tree.Node[T]) bool {
	parent := currentNode.Parent
	grandFather := parent.Parent
	return grandFather.Prev == parent && parent.Next == currentNode
}

func isRightRightCase[T comparable](currentNode *tree.Node[T]) bool {
	parent := currentNode.Parent
	grandFather := parent.Parent
	return grandFather.Next == parent && parent.Next == currentNode
}

func isRightLeftCase[T comparable](currentNode *tree.Node[T]) bool {
	parent := currentNode.Parent
	grandFather := parent.Parent
	return grandFather.Next == parent && parent.Prev == currentNode
}

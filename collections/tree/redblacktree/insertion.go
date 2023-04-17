package redblacktree

import (
	"collections/common"
	"collections/tree"
	"collections/tree/binarytree"
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
	// TODO
}

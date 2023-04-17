package rotation

import "collections/tree"

func RotateToRight[T comparable](nodeObj *tree.Node[T]) {
	if nodeObj.Parent == nil {
		return
	}
	nodeObjNext := nodeObj.Next
	parent := nodeObj.Parent
	nodeObj.Next = parent
	parent.Prev = nodeObjNext
	nodeObj.Parent = parent.Parent
	parent.Parent = nodeObj
	if nodeObj.Parent == nil {
		return
	}
	if nodeObj.Parent.Next == parent {
		nodeObj.Parent.Next = nodeObj
	} else {
		nodeObj.Parent.Prev = nodeObj
	}
}

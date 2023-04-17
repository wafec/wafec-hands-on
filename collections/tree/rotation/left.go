package rotation

import "collections/tree"

func RotateToLeft[T comparable](nodeObj *tree.Node[T]) {
	if nodeObj.Parent == nil {
		return
	}
	nodeObjPrev := nodeObj.Prev
	parent := nodeObj.Parent
	nodeObj.Prev = parent
	parent.Next = nodeObjPrev
	nodeObj.Parent = parent.Parent
	parent.Parent = nodeObj
	if nodeObj.Parent == nil {
		return
	} else if nodeObj.Parent.Next == parent {
		nodeObj.Parent.Next = nodeObj
	} else {
		nodeObj.Parent.Prev = nodeObj
	}
}

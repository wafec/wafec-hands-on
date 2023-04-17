package rotation

import "collections/tree"

func getTestNodes() (*tree.Node[int], *tree.Node[int], *tree.Node[int], *tree.Node[int], *tree.Node[int], *tree.Node[int], *tree.Node[int]) {
	obj10 := tree.NewNode(10, nil)
	obj8 := tree.NewNode(8, obj10)
	obj9 := tree.NewNode(9, obj8)
	obj7 := tree.NewNode(7, obj8)
	obj12 := tree.NewNode(12, obj10)
	obj11 := tree.NewNode(11, obj12)
	obj13 := tree.NewNode(13, obj12)
	obj10.Next = obj12
	obj10.Prev = obj8
	obj12.Next = obj13
	obj12.Prev = obj11
	obj8.Next = obj9
	obj8.Prev = obj7

	return obj7, obj8, obj9, obj10, obj11, obj12, obj13
}

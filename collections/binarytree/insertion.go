package binarytree

import "collections/common"

func Insert[T comparable](parentNode, valueNode Node[T], compareTo common.CompareTo[T]) {
	if parentNode != nil {
		if compareTo(valueNode.GetValue(), parentNode.GetValue()) > 0 {
			if parentNode.GetRightNode() == nil {
				parentNode.SetRightNode(valueNode)
				valueNode.SetParentNode(parentNode)
			} else {
				Insert(parentNode.GetRightNode(), valueNode, compareTo)
			}
		} else {
			if parentNode.GetLeftNode() == nil {
				parentNode.SetLeftNode(valueNode)
				valueNode.SetParentNode(parentNode)
			} else {
				Insert(parentNode.GetLeftNode(), valueNode, compareTo)
			}
		}
	}
}

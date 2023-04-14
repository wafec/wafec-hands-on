package binarytree

import "collections/common"

func Insert[T comparable](parentNode, valueNode Node[T], compareTo common.CompareTo[T]) Node[T] {
	if parentNode == nil {
		return valueNode
	} else {
		if compareTo(valueNode.GetValue(), parentNode.GetValue()) > 0 {
			if parentNode.GetRightNode() == nil {
				parentNode.SetRightNode(valueNode)
				valueNode.SetParentNode(parentNode)
				return valueNode
			} else {
				return Insert(parentNode.GetRightNode(), valueNode, compareTo)
			}
		} else {
			if parentNode.GetLeftNode() == nil {
				parentNode.SetLeftNode(valueNode)
				valueNode.SetParentNode(parentNode)
				return valueNode
			} else {
				return Insert(parentNode.GetLeftNode(), valueNode, compareTo)
			}
		}
	}
}

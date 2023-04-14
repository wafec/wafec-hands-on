package binarytree

type Node[T comparable] interface {
	GetValue() T
	SetValue(value T)
	GetLeftNode() Node[T]
	SetLeftNode(leftNode Node[T])
	GetRightNode() Node[T]
	SetRightNode(rightNode Node[T])
	GetParentNode() Node[T]
	SetParentNode(parentNode Node[T])
}

type NodeImpl[T comparable] struct {
	Value      T
	LeftNode   Node[T]
	RightNode  Node[T]
	ParentNode Node[T]
}

func (self *NodeImpl[T]) GetValue() T {
	return self.Value
}

func (self *NodeImpl[T]) SetValue(value T) {
	self.Value = value
}

func (self *NodeImpl[T]) GetLeftNode() Node[T] {
	return self.LeftNode
}

func (self *NodeImpl[T]) SetLeftNode(leftNode Node[T]) {
	self.LeftNode = leftNode
}

func (self *NodeImpl[T]) GetRightNode() Node[T] {
	return self.RightNode
}

func (self *NodeImpl[T]) SetRightNode(rightNode Node[T]) {
	self.RightNode = rightNode
}

func (self *NodeImpl[T]) GetParentNode() Node[T] {
	return self.ParentNode
}

func (self *NodeImpl[T]) SetParentNode(parentNode Node[T]) {
	self.ParentNode = parentNode
}

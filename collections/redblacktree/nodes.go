package redblacktree

import "collections/binarytree"

type Color int8

const (
	Red   Color = 1
	Black Color = 2
)

type Node[T comparable] interface {
	GetColor() Color
	SetColor(color Color)
	binarytree.Node[T]
}

type NodeImpl[T comparable] struct {
	Color Color
	Node[T]
}

func (self *NodeImpl[T]) GetColor() Color {
	return self.Color
}

func (self *NodeImpl[T]) SetColor(color Color) {
	self.Color = color
}

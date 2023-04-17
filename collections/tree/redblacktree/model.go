package redblacktree

import "collections/tree"

type Color int8

const (
	Red   Color = 1
	Black Color = 2
)

type Meta[T comparable] interface {
	GetColor(node *tree.Node[T]) Color
	SetColor(node *tree.Node[T], color Color)
}

type MetaMap[T comparable] struct {
	Map map[*tree.Node[T]]Color
}

func (self *MetaMap[T]) GetColor(node *tree.Node[T]) Color {
	color, ok := self.Map[node]
	if ok {
		return color
	} else {
		self.SetColor(node, Red)
		return Black
	}
}

func (self *MetaMap[T]) SetColor(node *tree.Node[T], color Color) {
	self.Map[node] = color
}

func NewMeta[T comparable]() Meta[T] {
	return &MetaMap[T]{Map: make(map[*tree.Node[T]]Color)}
}

package tree

type Node[T comparable] struct {
	Value  T
	Next   *Node[T]
	Prev   *Node[T]
	Parent *Node[T]
}

func NewNode[T comparable](value T, parent *Node[T]) *Node[T] {
	return &Node[T]{Value: value, Parent: parent}
}

func (self *Node[T]) IsLeaf() bool {
	return self.Next == nil && self.Prev == nil
}

func (self *Node[T]) GetUncleOrNil() *Node[T] {
	if self.Parent != nil && self.Parent.Parent != nil {
		grand := self.Parent.Parent
		if grand.Prev == self.Parent {
			return grand.Next
		} else {
			return grand.Prev
		}
	} else {
		return nil
	}
}

func (self *Node[T]) GetGrandFatherOrNil() *Node[T] {
	if self.Parent != nil && self.Parent.Parent != nil {
		return self.Parent.Parent
	} else {
		return nil
	}
}

type Tree[T comparable] struct {
	Root *Node[T]
}

package binarytree

import (
	"collections/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert_ParentIsNull(t *testing.T) {
	assert := assert.New(t)
	var input Node[int64] = &NodeImpl[int64]{Value: 10}

	Insert(nil, input, common.Int64CompareTo)

	assert.Nil(input.GetLeftNode())
	assert.Nil(input.GetRightNode())
	assert.Nil(input.GetParentNode())
}

func TestInsert_ValueIsGreaterThanRootValue(t *testing.T) {
	assert := assert.New(t)
	var input Node[int64] = &NodeImpl[int64]{Value: 10}
	var rootNode Node[int64] = &NodeImpl[int64]{Value: 9}

	Insert(rootNode, input, common.Int64CompareTo)

	assert.Equal(rootNode.GetRightNode().GetValue(), input.GetValue(), "Root right node value should be equal to input value")
	assert.Equal(input.GetParentNode().GetValue(), rootNode.GetValue(), "Input parent node value should be equal to root value")
	assert.Nil(rootNode.GetLeftNode())
}

func TestInsert_ValueIsLesserThanRootValue(t *testing.T) {
	assert := assert.New(t)
	var input Node[int64] = &NodeImpl[int64]{Value: 10}
	var rootNode Node[int64] = &NodeImpl[int64]{Value: 11}

	Insert(rootNode, input, common.Int64CompareTo)

	assert.Equal(rootNode.GetLeftNode().GetValue(), input.GetValue(), "Root left node value should be equal to input value")
	assert.Equal(input.GetParentNode().GetValue(), rootNode.GetValue(), "Input parent node value should be equal to root value")
	assert.Nil(rootNode.GetRightNode())
}

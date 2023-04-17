package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLeaf_True(t *testing.T) {
	assert := assert.New(t)
	nodeObj := NewNode(10, nil)

	valueReturned := nodeObj.IsLeaf()

	assert.True(valueReturned)
}

func TestIsLeaf_False(t *testing.T) {
	assert := assert.New(t)
	nodeObj := NewNode(10, nil)
	nodeObj.Next = NewNode(11, nil)

	valueReturned := nodeObj.IsLeaf()

	assert.False(valueReturned)
}

func TestGetUncleOrNil_NoParent(t *testing.T) {
	assert := assert.New(t)
	nodeObj := NewNode(10, nil)

	result := nodeObj.GetUncleOrNil()

	assert.Nil(result)
}

func TestGetUncleOrNil_NoGrand(t *testing.T) {
	assert := assert.New(t)
	nodeObj := NewNode(9, nil)
	nodeObj.Parent = NewNode(10, nil)
	nodeObj.Parent.Prev = nodeObj

	result := nodeObj.GetUncleOrNil()

	assert.Nil(result)
}

func TestGetUncleOrNil_NoUncle(t *testing.T) {
	assert := assert.New(t)
	nodeObj := NewNode(9, nil)
	nodeObj.Parent = NewNode(10, nil)
	nodeObj.Parent.Prev = nodeObj
	nodeObj.Parent.Parent = NewNode(11, nil)
	nodeObj.Parent.Parent.Prev = nodeObj.Parent

	result := nodeObj.GetUncleOrNil()

	assert.Nil(result)
}

func TestGetUncleOrNil_UncleReturned(t *testing.T) {
	assert := assert.New(t)
	nodeObj := NewNode(9, nil)
	nodeObj.Parent = NewNode(10, nil)
	nodeObj.Parent.Parent = NewNode(11, nil)
	nodeObj.Parent.Parent.Next = NewNode(12, nil)
	nodeObj.Parent.Prev = nodeObj
	nodeObj.Parent.Parent.Prev = nodeObj.Parent
	nodeObj.Parent.Parent.Next.Parent = nodeObj.Parent.Parent

	uncleReturned := nodeObj.GetUncleOrNil()

	assert.Equal(nodeObj.Parent.Parent.Next, uncleReturned)
}

func TestGetGrandFatherOrNil_NoParent(t *testing.T) {
	assert := assert.New(t)
	nodeObj := NewNode(10, nil)

	result := nodeObj.GetGrandFatherOrNil()

	assert.Nil(result)
}

func TestGetGrandFatherOrNil_NoGrandFather(t *testing.T) {
	assert := assert.New(t)
	parentObj := NewNode(11, nil)
	nodeObj := NewNode(10, parentObj)
	parentObj.Prev = nodeObj

	result := nodeObj.GetGrandFatherOrNil()

	assert.Nil(result)
}

func TestGetGrandFatherOrNil_GrandFatherReturned(t *testing.T) {
	assert := assert.New(t)
	grandFather := NewNode(12, nil)
	parent := NewNode(11, grandFather)
	grandFather.Prev = parent
	child := NewNode(10, parent)
	parent.Prev = child

	result := child.GetGrandFatherOrNil()

	assert.Equal(grandFather, result)
}

package rotation

import (
	"collections/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateToRight_NoGrandFather(t *testing.T) {
	assert := assert.New(t)
	obj7, obj8, obj9, obj10, obj11, obj12, obj13 := getTestNodes()

	RotateToRight(obj8)

	assert.Equal(obj10, obj8.Next)
	assert.Nil(obj8.Parent)
	assert.Equal(obj7, obj8.Prev)
	assert.Equal(obj10.Prev, obj9)
	assert.Equal(obj12, obj10.Next)
	assert.Equal(obj11, obj12.Prev)
	assert.Equal(obj13, obj12.Next)
}

func TestRotateToRight_GrandFatherOnRight(t *testing.T) {
	assert := assert.New(t)
	_, obj8, _, obj10, _, _, _ := getTestNodes()
	grandFather := tree.NewNode(50, nil)
	obj10.Parent = grandFather
	grandFather.Prev = obj10

	RotateToRight(obj8)

	assert.Equal(obj8, grandFather.Prev)
	assert.Equal(grandFather, obj8.Parent)
}

func TestRotateToRight_GrandFatherOnLeft(t *testing.T) {
	assert := assert.New(t)
	_, obj8, _, obj10, _, _, _ := getTestNodes()
	grandFather := tree.NewNode(0, nil)
	grandFather.Next = obj10
	obj10.Parent = grandFather

	RotateToRight(obj8)

	assert.Equal(obj8, grandFather.Next)
	assert.Equal(grandFather, obj8.Parent)
}

package rotation

import (
	"collections/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateToLeft_NoGrandFather(t *testing.T) {
	assert := assert.New(t)
	obj7, obj8, obj9, obj10, obj11, obj12, obj13 := getTestNodes()

	RotateToLeft(obj12)

	assert.Equal(obj10, obj12.Prev)
	assert.Equal(obj11, obj10.Next)
	assert.Equal(obj8, obj10.Prev)
	assert.Equal(obj13, obj12.Next)
	assert.Equal(obj9, obj8.Next)
	assert.Equal(obj7, obj8.Prev)
}

func TestRotateToLeft_GrandFatherOnRight(t *testing.T) {
	assert := assert.New(t)
	_, _, _, obj10, _, obj12, _ := getTestNodes()
	grandFather := tree.NewNode(50, nil)
	grandFather.Prev = obj10
	obj10.Parent = grandFather

	RotateToLeft(obj12)

	assert.Equal(obj12, grandFather.Prev)
	assert.Equal(grandFather, obj12.Parent)
}

func TestRotateToLeft_GrandFatherOnLeft(t *testing.T) {
	assert := assert.New(t)
	_, _, _, obj10, _, obj12, _ := getTestNodes()
	grandFather := tree.NewNode(0, nil)
	grandFather.Next = obj10
	obj10.Parent = grandFather

	RotateToLeft(obj12)

	assert.Equal(obj12, grandFather.Next)
	assert.Equal(grandFather, obj12.Parent)
}

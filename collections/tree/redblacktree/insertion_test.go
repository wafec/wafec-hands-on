package redblacktree

import (
	"collections/common/comparators"
	"collections/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert_Empty(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int64]{}
	metaObj := NewMeta[int64]()

	insertedNode := Insert(&treeObj, 10, comparators.Int64CompareFunc, metaObj)

	assert.Equal(Black, metaObj.GetColor(insertedNode))
	assert.Equal(insertedNode, treeObj.Root)
}

func TestInsert_RightRightCase(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int]{}
	metaObj := NewMeta[int]()

	obj3 := Insert(&treeObj, 3, comparators.IntCompareFunc, metaObj)
	obj21 := Insert(&treeObj, 21, comparators.IntCompareFunc, metaObj)
	obj32 := Insert(&treeObj, 32, comparators.IntCompareFunc, metaObj)

	assert.Equal(obj3, obj21.Prev)
	assert.Equal(obj32, obj21.Next)
	assert.Equal(Red, metaObj.GetColor(obj3))
	assert.Equal(Red, metaObj.GetColor(obj32))
	assert.Equal(Black, metaObj.GetColor(obj21))
	assert.Equal(obj21, treeObj.Root)
}

func TestInsert_RightRightCasePlusParentRecoloring(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int]{}
	metaObj := NewMeta[int]()

	obj3 := Insert(&treeObj, 3, comparators.IntCompareFunc, metaObj)
	obj21 := Insert(&treeObj, 21, comparators.IntCompareFunc, metaObj)
	obj32 := Insert(&treeObj, 32, comparators.IntCompareFunc, metaObj)
	obj15 := Insert(&treeObj, 15, comparators.IntCompareFunc, metaObj)

	assert.Equal(Red, metaObj.GetColor(obj15))
	assert.Equal(Black, metaObj.GetColor(obj3))
	assert.Equal(Black, metaObj.GetColor(obj32))
	assert.Equal(Black, metaObj.GetColor(obj21))
	assert.Equal(obj21, treeObj.Root)
}

func TestInsert_LeftLeftCase(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int]{}
	metaObj := NewMeta[int]()

	obj32 := Insert(&treeObj, 32, comparators.IntCompareFunc, metaObj)
	obj21 := Insert(&treeObj, 21, comparators.IntCompareFunc, metaObj)
	obj3 := Insert(&treeObj, 3, comparators.IntCompareFunc, metaObj)

	assert.Equal(obj3, obj21.Prev)
	assert.Equal(obj32, obj21.Next)
	assert.Equal(Red, metaObj.GetColor(obj3))
	assert.Equal(Red, metaObj.GetColor(obj32))
	assert.Equal(Black, metaObj.GetColor(obj21))
	assert.Equal(obj21, treeObj.Root)
}

func TestInsert_LeftLeftCasePlusParentRecoloring(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int]{}
	metaObj := NewMeta[int]()

	obj32 := Insert(&treeObj, 32, comparators.IntCompareFunc, metaObj)
	obj21 := Insert(&treeObj, 21, comparators.IntCompareFunc, metaObj)
	obj3 := Insert(&treeObj, 3, comparators.IntCompareFunc, metaObj)
	obj29 := Insert(&treeObj, 29, comparators.IntCompareFunc, metaObj)

	assert.Equal(Black, metaObj.GetColor(obj21))
	assert.Equal(Black, metaObj.GetColor(obj3))
	assert.Equal(Black, metaObj.GetColor(obj32))
	assert.Equal(Red, metaObj.GetColor(obj29))
	assert.Equal(obj21, treeObj.Root)
}

func TestInsert_RightLeftCase(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int]{}
	metaObj := NewMeta[int]()

	obj1 := Insert(&treeObj, 1, comparators.IntCompareFunc, metaObj)
	obj2 := Insert(&treeObj, 2, comparators.IntCompareFunc, metaObj)
	obj3 := Insert(&treeObj, 3, comparators.IntCompareFunc, metaObj)
	obj5 := Insert(&treeObj, 5, comparators.IntCompareFunc, metaObj)
	obj4 := Insert(&treeObj, 4, comparators.IntCompareFunc, metaObj)

	assert.Equal(obj2, treeObj.Root)
	assert.Equal(Red, metaObj.GetColor(obj3))
	assert.Equal(Red, metaObj.GetColor(obj5))
	assert.Equal(Black, metaObj.GetColor(obj4))
	assert.Equal(Black, metaObj.GetColor(obj2))
	assert.Equal(Black, metaObj.GetColor(obj1))
}

func TestInsert_LeftRightCase(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int]{}
	metaObj := NewMeta[int]()

	obj9 := Insert(&treeObj, 9, comparators.IntCompareFunc, metaObj)
	obj8 := Insert(&treeObj, 8, comparators.IntCompareFunc, metaObj)
	obj5 := Insert(&treeObj, 5, comparators.IntCompareFunc, metaObj)
	obj7 := Insert(&treeObj, 7, comparators.IntCompareFunc, metaObj)
	obj6 := Insert(&treeObj, 6, comparators.IntCompareFunc, metaObj)

	assert.Equal(obj8, treeObj.Root)
	assert.Equal(Black, metaObj.GetColor(obj8))
	assert.Equal(Black, metaObj.GetColor(obj6))
	assert.Equal(Black, metaObj.GetColor(obj9))
	assert.Equal(Red, metaObj.GetColor(obj5))
	assert.Equal(Red, metaObj.GetColor(obj7))

}

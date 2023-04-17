package binarytree

import (
	"collections/common/comparators"
	"collections/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert_Empty(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int64]{}

	nodeInserted := Insert(&treeObj, 10, comparators.Int64CompareFunc)

	assert.NotNil(nodeInserted)
	assert.NotNil(treeObj.Root)
	assert.Equal(treeObj.Root, nodeInserted)
	assert.True(treeObj.Root.IsLeaf())
}

func TestInsert_AddedToPrevious(t *testing.T) {
	assert := assert.New(t)
	treeObj := tree.Tree[int64]{}
	treeObj.Root = tree.NewNode[int64](10, nil)

	nodeInserted := Insert(&treeObj, 9, comparators.Int64CompareFunc)

	assert.Equal(treeObj.Root, nodeInserted.Parent)
	assert.Equal(nodeInserted, treeObj.Root.Prev)
	assert.Equal(int64(9), nodeInserted.Value)
}

func TestInsert_AddedToNext(t *testing.T) {

}

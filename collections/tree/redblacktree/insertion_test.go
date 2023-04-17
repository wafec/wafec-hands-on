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
}

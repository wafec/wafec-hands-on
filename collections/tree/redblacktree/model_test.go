package redblacktree

import (
	"collections/tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetColor_Empty(t *testing.T) {
	assert := assert.New(t)
	nodeObj := tree.NewNode[int64](10, nil)
	metaObj := NewMeta[int64]()

	metaObj.SetColor(nodeObj, Red)

	assert.Equal(Red, metaObj.GetColor(nodeObj))
}

func TestGetColor_Defaults(t *testing.T) {
	assert := assert.New(t)
	nodeObj := tree.NewNode[int64](10, nil)
	metaObj := NewMeta[int64]()

	colorReturned := metaObj.GetColor(nodeObj)

	assert.Equal(Black, colorReturned)
}

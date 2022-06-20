
package alat

import (
	"syscall/js"
)

type Block struct {
	htmlObject js.Value
	widgets []Widget
}

func NewBlock(mainHtmlObject js.Value) *Block {
	var block Block
	block.htmlObject = mainHtmlObject
	return &block
}

func (b *Block) Draw() {
	for _,widget := range b.widgets {
		widget.Draw()
	}
}
func (b *Block) Refresh() {}


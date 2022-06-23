
package alat

import (
	"syscall/js"
)

type Block struct {
	htmlObject js.Value
	widgets []Widget
	buffer *Buffer
	columnsToWidgets map[string]Widget
	widgetsToColumns map[Widget]string
}

func NewBlock(mainHtmlObject js.Value) *Block {
	var block Block
	block.htmlObject = mainHtmlObject
	block.buffer = NewBuffer()
	block.columnsToWidgets = make(map[string]Widget)
	block.widgetsToColumns = make(map[Widget]string)
	return &block
}

func (b *Block) Buffer() *Buffer {
	return b.buffer
}

func (b *Block) Connect(widget Widget, column string) {
	b.columnsToWidgets[column] = widget
	b.widgetsToColumns[widget] = column
}


func (b *Block) Draw() {
	for _,widget := range b.widgets {
		widget.Draw()
	}
}
func (b *Block) Refresh() {}


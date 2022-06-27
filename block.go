
package alat

import (
	"syscall/js"
	"fmt"
)

type Block struct {
	htmlObject js.Value
	widgets []Widget
	buffer *Buffer
	columnsToWidgets map[string]Widget
	widgetsToColumns map[Widget]string
	focusList []Widget
	focusIndexed map[Widget]int
	currentWidget Widget
}

func NewBlock(mainHtmlObject js.Value) *Block {
	var block Block
	block.htmlObject = mainHtmlObject
	block.buffer = NewBuffer()
	block.columnsToWidgets = make(map[string]Widget)
	block.widgetsToColumns = make(map[Widget]string)
	block.focusList = make([]Widget,0)
	block.focusIndexed = make(map[Widget]int)
	block.currentWidget = nil
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
	b.FirstWidget().SetFocus()
}


func (b *Block) Refresh() {
	for _,widget := range b.widgets {
		widget.Refresh()
	}
}


func (b *Block) AddToFocusList(widget Widget) {
	b.focusList = append(b.focusList,widget)
	b.focusIndexed[widget] = len(b.focusList)-1
}


func (b *Block) OnFocusToWidget(widget Widget) {
	b.currentWidget = widget
	fmt.Println("Focus set to column:",b.widgetsToColumns[b.currentWidget])
}


func (b *Block) FocusCurrent() {
	if b.currentWidget != nil {
		b.currentWidget.SetFocus()
	}
}


func (b *Block) FirstWidget() Widget { 
	return b.focusList[0]
}


func (b *Block) LastWidget() Widget { 
	return b.focusList[len(b.focusList)-1]
}


func (b *Block) NextWidget() Widget { 
	if b.currentWidget == nil {
		return b.FirstWidget()
	}
	idx := b.focusIndexed[b.currentWidget]+1
	if idx >= len(b.focusList) {
		return b.FirstWidget()
	}
	return b.focusList[idx]
}


func (b *Block) PrevWidget() Widget { 
	if b.currentWidget == nil {
		return b.FirstWidget()
	}
	idx := b.focusIndexed[b.currentWidget]-1
	if idx < 0 {
		return b.LastWidget()
	}
	return b.focusList[idx]
}





package alat

import (
	"syscall/js"
	"fmt"
)

type Block struct {
	htmlObject js.Value
	widgets []Widget
	buffer *Buffer
	columnsToWidgets map[string]FocusableWidget
	widgetsToColumns map[FocusableWidget]string
	focusList []FocusableWidget
	focusIndexed map[FocusableWidget]int
	currentWidget FocusableWidget
	viewBegin int
	viewEnd int
	visibleRows int
	lastFocusOutWidget FocusableWidget
	lastFocusOutPos int
}

func NewBlock(mainHtmlObject js.Value, visibleRows int) *Block {
	var block Block
	block.htmlObject = mainHtmlObject
	block.buffer = NewBuffer()
	block.columnsToWidgets = make(map[string]FocusableWidget)
	block.widgetsToColumns = make(map[FocusableWidget]string)
	block.focusList = make([]FocusableWidget,0)
	block.focusIndexed = make(map[FocusableWidget]int)
	block.currentWidget = nil
	block.visibleRows = visibleRows
	return &block
}

func (b *Block) Buffer() *Buffer {
	return b.buffer
}

func (b *Block) Connect(widget FocusableWidget, column string) {
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
	//fmt.Println("Block Refresh")
	for _,widget := range b.widgets {
		widget.Refresh()
	}
}

func (b *Block) RefreshCurrentRow() {
	//fmt.Println("Block Refresh Current Row")
	for _,widget := range b.widgets {
		widget.RefreshCurrentRow()
	}
}


func (b *Block) AddToFocusList(widget FocusableWidget) {
	b.focusList = append(b.focusList,widget)
	b.focusIndexed[widget] = len(b.focusList)-1
}


func (b *Block) FocusCurrent() {
	if b.currentWidget != nil {
		b.currentWidget.SetFocus()
	}
}


func (b *Block) FirstWidget() FocusableWidget { 
	return b.focusList[0]
}


func (b *Block) LastWidget() FocusableWidget { 
	return b.focusList[len(b.focusList)-1]
}


func (b *Block) NextWidget() FocusableWidget { 
	if b.currentWidget == nil {
		return b.FirstWidget()
	}
	idx := b.focusIndexed[b.currentWidget]+1
	if idx >= len(b.focusList) {
		return b.FirstWidget()
	}
	return b.focusList[idx]
}


func (b *Block) PrevWidget() FocusableWidget { 
	if b.currentWidget == nil {
		return b.FirstWidget()
	}
	idx := b.focusIndexed[b.currentWidget]-1
	if idx < 0 {
		return b.LastWidget()
	}
	return b.focusList[idx]
}


func (b *Block) Pos() int { 
	return b.buffer.pos
}

func (b *Block) ViewRow2BufferPos(viewRow int) (int, bool) {
	if viewRow>=0 && viewRow<=b.visibleRows {
		bufferPos := viewRow+b.viewBegin
		if bufferPos>=0 && bufferPos<len(b.buffer.rows) {
			return bufferPos,true
		}
	}
	return 0,false
}

func (b *Block) BufferPos2ViewRow(bufferPos int) (int, bool) {
	if bufferPos>=0 && bufferPos<len(b.buffer.rows) {
		viewRow := bufferPos-b.viewBegin
		if viewRow>=0 && viewRow<b.visibleRows {
			return viewRow,true
		}
	}
	return 0,false
}

func (b *Block) RownumState(rownum int) int {
	bufferRow := b.viewBegin+rownum
	pos := b.Buffer().pos
	if bufferRow>=0 && bufferRow<=b.viewEnd {
		if bufferRow==pos {
			return ROWNUM_CURRENT
		}
		return ROWNUM_NOT_CURRENT
	}
	return ROWNUM_NOT_EXISTS
}


func (b *Block) OnFocusToWidget(widget FocusableWidget) {
	b.currentWidget = widget
	//fmt.Println("Focus set to column:",b.widgetsToColumns[b.currentWidget])
}


func (b *Block) GotoRequest(newWidget FocusableWidget, newPos int) bool {
	lastWidget := b.lastFocusOutWidget
	lastPos := b.lastFocusOutPos
	response := false
	fmt.Printf("Goto Request from %p:%d to %p:%d = %v",
		lastWidget,lastPos,newWidget,newPos,response)
	return response
}


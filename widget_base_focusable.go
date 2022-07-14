
package alat

import (
	"syscall/js"
)

type FocusableWidget interface {
	Widget
	SetFocus()	
	WriteIfChanged(obj js.Value, rownum int) bool
	SelectAll()
	Label() string
	SetLabel(string)
	Index() int
	SetIndex(int)
	IsInMultiRow() bool
	DrawInMultiRow(js.Value, int) js.Value 
	ColumnCell(js.Value) js.Value
	RefreshCell(js.Value, int, string)
}

type BaseFocusableWidget struct {
	BaseWidget
	label string
	index int
}

func (w *BaseFocusableWidget) Init(block *Block, widget FocusableWidget, parentWidget ParentWidget, label string) {
	w.BaseWidget.Init(block, widget, parentWidget)
	block.AddToFocusList(widget)
	if w.IsInMultiRow() {
		w.index = len(parentWidget.(ParentWidget).Children())-1
	}
	w.label = label
}

func (w *BaseFocusableWidget) SetFocus() {
	w.HTMLObject().Call("focus")
}

func (w *BaseFocusableWidget) SelectAll() {
	w.HTMLObject().Call("select")
}

func (w *BaseFocusableWidget) WriteIfChanged(obj js.Value, rownum int) bool {
	return false
}

func (w *BaseFocusableWidget) IsFocusable() bool {
	return true
}

func (w *BaseFocusableWidget) Label() string {
	return w.label
}
	
func (w *BaseFocusableWidget) SetLabel(label string) {
	w.label = label
}

func (w *BaseFocusableWidget) Index() int {
	return w.index
}
	
func (w *BaseFocusableWidget) SetIndex(index int) {
	w.index = index
}

func (w *BaseFocusableWidget) IsInMultiRow() bool {
	return w.ParentWidget().IsMultiRow()
}

func (w *BaseFocusableWidget) DrawInMultiRow(parent js.Value, rownum int) js.Value {
	input := NewNode(parent,"input")
	AttachFocusEvents(w,input,rownum)
	return input
}

func (w *BaseFocusableWidget) RefreshCell(cell js.Value, rowState int, value string) {
	switch rowState {
	case ROWNUM_NOT_EXISTS:
		cell.Set("style","background-color: #CCC")
		cell.Set("type","hidden")
	case ROWNUM_NOT_CURRENT:
		cell.Set("style","background-color: #FFFFFF")
		cell.Set("type","")
		cell.Set("value",value)
	case ROWNUM_CURRENT:
		cell.Set("style","background-color: #BCE5FD")
		cell.Set("type","")
		cell.Set("value",value)
	}
}

func (w *BaseFocusableWidget) HTMLObject() js.Value {
	if w.IsInMultiRow() {
		table := w.ParentWidget().(MultiRowWidget)
		buffer := w.Block().Buffer()
		viewRow,_ := w.Block().BufferPos2ViewRow(buffer.pos)
		tdElement := table.CellElement(w.index,viewRow) 
		return w.ColumnCell(tdElement) // Beware for overriding!
	}
	return w.BaseWidget.HTMLObject()
}

// If overriden, HTMLObject should be too
func (w *BaseFocusableWidget) ColumnCell(parent js.Value) js.Value {
	return parent.Get("children").Get("0")
}



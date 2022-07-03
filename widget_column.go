package alat

import (
	"syscall/js"
	"fmt"
)

type Column struct {
	BaseWidget
	label string
	index int
}

func (w *Column) HTMLObject() js.Value {
	table := w.ParentWidget().(*Table)
	buffer := w.Block().Buffer()
	viewRow,_ := w.Block().BufferPos2ViewRow(buffer.pos)
	return table.Cell(w.index,viewRow)
}

func (w *Column) SetFocus() {
	w.HTMLObject().Call("focus")
}

func (w *Column) SelectAll() {
	w.HTMLObject().Call("select")
}


func (w *Column) WriteIfChanged(obj js.Value, rownum int) bool {
	if column,ok := w.Block().widgetsToColumns[w]; ok {
		pos,_ := w.Block().ViewRow2BufferPos(rownum)
		buffer := w.Block().Buffer()
		widgetValue := obj.Get("value").String()
		bufferValue,_ := buffer.GetAt(pos,column)
		if widgetValue != bufferValue {
			fmt.Println("New value for",column,"at",pos,"is",widgetValue)
			buffer.SetAt(pos,column,widgetValue)
			return true
		}
	}
	return false
}
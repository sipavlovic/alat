package alat

import (
	"syscall/js"
)

type Column struct {
	BaseWidget
	label string
	index int
}

func (w *Column) HTMLObject() js.Value {
	table := w.ParentWidget().(*Table)
	buffer := w.Block().Buffer()
	viewRow,_ := table.BufferPos2ViewRow(buffer.pos)
	return table.Cell(w.index,viewRow)
}
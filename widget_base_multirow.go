
package alat

import (
	"syscall/js"
)

type MultiRowWidget interface {
	ParentWidget
	CellElement(int, int) js.Value
}

type BaseMultiRowWidget struct {
	BaseParentWidget
}

func (w *BaseMultiRowWidget) Init(block *Block, widget Widget, parentWidget ParentWidget) {
	w.BaseParentWidget.Init(block, widget, parentWidget)
}

func (w *BaseMultiRowWidget) IsMultiRow() bool {
	return true
}

func (w *BaseMultiRowWidget) CellElement(colnum int, rownum int) js.Value {
	panic("CellElement: TO OVERRIDE")
	return js.Null()
}




package alat

import (
	"syscall/js"
	"fmt"
)


type Edit struct {
	BaseWidget
}

func NewEdit(block *Block, parentWidget Widget) *Edit {
	var edit Edit
	edit.BaseWidget.Init(block, &edit, parentWidget)
	block.AddToFocusList(&edit)
	return &edit
}

func (w *Edit) Draw() {
	input := NewNode(w.ParentHTMLObject(),"input")
	w.htmlObject = input
	w.BaseWidget.Draw()
	AttachFocusEvents(w,input,NOTINTABLE)
}

func (w *Edit) Refresh() {
	result := ""
	if column,ok := w.Block().widgetsToColumns[w]; ok {
		result,_ = w.Block().Buffer().Get(column)
	}
	w.HTMLObject().Set("value",result)
	w.BaseWidget.Refresh()
}

func (w *Edit) RefreshCurrentRow() {
	result := ""
	if column,ok := w.Block().widgetsToColumns[w]; ok {
		result,_ = w.Block().Buffer().Get(column)
	}
	w.HTMLObject().Set("value",result)
	w.BaseWidget.RefreshCurrentRow()
}

func (w *Edit) SetFocus() {
	w.HTMLObject().Call("focus")
}

func (w *Edit) SelectAll() {
	w.HTMLObject().Call("select")
}


func (w *Edit) WriteIfChanged(obj js.Value, rownum int) bool {
	if column,ok := w.Block().widgetsToColumns[w]; ok {
		buffer := w.Block().Buffer()
		widgetValue := obj.Get("value").String()
		bufferValue,_ := buffer.Get(column)
		if widgetValue != bufferValue {
			fmt.Println("New value for",column,"is",widgetValue)
			buffer.Set(column,widgetValue)
			return true
		}
	}
	return false
}

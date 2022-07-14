
package alat

import (
	"syscall/js"
	"fmt"
)


type Edit struct {
	BaseFocusableWidget
}

func NewEdit(block *Block, parentWidget ParentWidget, label string) *Edit {
	var edit Edit
	edit.BaseFocusableWidget.Init(block, &edit, parentWidget, label)
	return &edit
}

func (w *Edit) Draw() {
	divLabel := NewNode(w.ParentHTMLObject(),"div")
	divLabel.Set("textContent",w.Label())
	divInput := NewNode(w.ParentHTMLObject(),"div")
	input := NewNode(divInput,"input")
	w.htmlObject = input
	w.BaseFocusableWidget.Draw()
	AttachFocusEvents(w,input,NOTINTABLE)
}

func (w *Edit) DrawInMultiRow(parent js.Value, rownum int) js.Value {
	input := NewNode(parent,"input")
	AttachFocusEvents(w,input,rownum)
	return input
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

func (w *Edit) WriteIfChanged(obj js.Value, rownum int) bool {
	if column,ok := w.Block().widgetsToColumns[w]; ok {
		buffer := w.Block().Buffer()
		widgetValue := obj.Get("value").String()
		bufferValue := ""
		if w.IsInMultiRow() {
			pos,_ := w.Block().ViewRow2BufferPos(rownum)
			bufferValue,_ = buffer.GetAt(pos,column)
			if widgetValue != bufferValue {
				fmt.Println("New value for",column,"at",rownum,"is",widgetValue)
				buffer.SetAt(pos,column,widgetValue)
				return true
			}
		} else {
			bufferValue,_ = buffer.Get(column)
			if widgetValue != bufferValue {
				fmt.Println("New value for",column,"is",widgetValue)
				buffer.Set(column,widgetValue)
				return true
			}
		}
	}
	return false
}


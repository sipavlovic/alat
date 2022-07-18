
package alat

import (
	"syscall/js"
)

type Widget interface {
	Block() *Block
	ParentWidget() Widget
	ParentHTMLObject() js.Value
	HTMLObject() js.Value
	SetHTMLObject(js.Value)
	Draw()
	Refresh()
	RefreshCurrentRow()
	Remove()

	IsFocusable() bool
	IsParent() bool
	IsMultiRow() bool
}





type BaseWidget struct {
	block *Block
	parentWidget ParentWidget
	//children []Widget
	htmlParent js.Value
	htmlObject js.Value
}

func (w *BaseWidget) Init(block *Block, widget Widget, parentWidget ParentWidget) {
	w.block = block
	w.parentWidget = parentWidget
	if w.parentWidget != nil {
		w.parentWidget.AddChild(widget)	
	} else {
		w.block.widgets = append(w.block.widgets,widget)
	}
}

func (w *BaseWidget) Block() *Block { 
	return w.block
}

func (w *BaseWidget) ParentWidget() Widget {
	return w.parentWidget
}

func (w *BaseWidget) ParentHTMLObject() js.Value {
	if w.parentWidget != nil {
		return w.parentWidget.HTMLObject()
	} 
	return w.Block().htmlObject
}

func (w *BaseWidget) HTMLObject() js.Value { 
	return w.htmlObject 
}

func (w *BaseWidget) SetHTMLObject(obj js.Value) {
	w.htmlObject = obj
}

func (w *BaseWidget) Draw() {}

func (w *BaseWidget) Refresh() {}

func (w *BaseWidget) RefreshCurrentRow() {}

func (w *BaseWidget) Remove() {}

func (w *BaseWidget) IsFocusable() bool {
	return false
}

func (w *BaseWidget) IsParent() bool {
	return false
}

func (w *BaseWidget) IsMultiRow() bool {
	return false
}





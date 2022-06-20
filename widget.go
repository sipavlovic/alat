
package alat

import (
	"syscall/js"
)

type Widget interface {
	Block() *Block
	ParentWidget() Widget
	ParentHTMLObject() js.Value
	HTMLObject() js.Value
	Draw()
	Refresh()
	AddChild(Widget)
	Children() []Widget 
	SetFocus()	
}


type BaseWidget struct {
	block *Block
	parentWidget Widget
	children []Widget
	htmlParent js.Value
	htmlObject js.Value
}

func (w *BaseWidget) Init(block *Block, widget Widget, parentWidget Widget) {
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

func (w *BaseWidget) Draw() {
	for _, child := range w.children {
		child.Draw()
	}
}

func (w *BaseWidget) Refresh() {
	for _, child := range w.children {
		child.Refresh()
	}
}

func (w *BaseWidget) AddChild(child Widget) {
	w.children = append(w.children,child)
}

func (w *BaseWidget) Children() []Widget {
	return w.children
}

func (w *BaseWidget) SetFocus()	{}



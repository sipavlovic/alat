
package alat

import (
	"syscall/js"
	"fmt"
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
	AddChild(Widget)
	Children() []Widget 
	SetFocus()	
	WriteIfChanged(obj js.Value, rownum int) bool
	SelectAll()
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

func (w *BaseWidget) RefreshCurrentRow() {
	for _, child := range w.children {
		child.RefreshCurrentRow()
	}
}


func (w *BaseWidget) AddChild(child Widget) {
	w.children = append(w.children,child)
}

func (w *BaseWidget) Children() []Widget {
	return w.children
}

func (w *BaseWidget) SetFocus()	{}

func (w *BaseWidget) WriteIfChanged(obj js.Value, rownum int) bool {
	return false
}

func (w *BaseWidget) SelectAll() {}

func (w *BaseWidget) SetHTMLObject(obj js.Value) {
	w.htmlObject = obj
}

// ----------------------------------------------------------


func AttachFocusEvents(widget Widget, obj js.Value, rownum int) {
	obj.Set("onkeydown",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		keycode := event.Get("keyCode").Int()
		shiftkey := event.Get("shiftKey").Bool()
		ctrlkey := event.Get("ctrlKey").Bool()
		//fmt.Println("OnKeyDown:",keycode,shiftkey,ctrlkey,rownum)
		switch keycode {
		case 9: // Tab
			event.Call("preventDefault")
			if shiftkey && !ctrlkey {
				widget.WriteIfChanged(obj,rownum)
				widget.Block().Refresh()
				widget.Block().PrevWidget().SetFocus()
			} else if !shiftkey && !ctrlkey {
				widget.WriteIfChanged(obj,rownum)
				widget.Block().Refresh()
				widget.Block().NextWidget().SetFocus()
			}
		case 38: // Up
			event.Call("preventDefault")
			if !shiftkey && !ctrlkey {
				widget.WriteIfChanged(obj,rownum)
				buffer := widget.Block().Buffer()
				lastPos := buffer.pos
				buffer.Goto(lastPos-1)
				if lastPos != buffer.pos {
					widget.Block().Refresh()
					widget.SetFocus()
				}
				widget.SelectAll()
			}
		case 40: // Down
			event.Call("preventDefault")
			if !shiftkey && !ctrlkey {
				widget.WriteIfChanged(obj,rownum)
				buffer := widget.Block().Buffer()
				lastPos := buffer.pos
				buffer.Goto(lastPos+1)
				if lastPos != buffer.pos {
					widget.Block().Refresh()
					widget.SetFocus()
				}
				widget.SelectAll()
			}
		case 33: // PgUp
			event.Call("preventDefault")
			if !shiftkey && !ctrlkey {
				widget.WriteIfChanged(obj,rownum)
				block := widget.Block()
				buffer := block.Buffer()
				lastPos := buffer.pos
				if lastPos == block.viewBegin {
					buffer.Goto(lastPos-block.visibleRows)
				} else {
					buffer.Goto(block.viewBegin)
				}
				if lastPos != buffer.pos {
					block.Refresh()
					widget.SetFocus()
				}
				widget.SelectAll()
			}
		case 34: // PgDn
			event.Call("preventDefault")
			if !shiftkey && !ctrlkey {
				widget.WriteIfChanged(obj,rownum)
				block := widget.Block()
				buffer := block.Buffer()
				lastPos := buffer.pos
				if lastPos == block.viewEnd {
					buffer.Goto(lastPos+block.visibleRows)
				} else {
					buffer.Goto(block.viewEnd)
				}
				if lastPos != buffer.pos {
					block.Refresh()
					widget.SetFocus()
				}
				widget.SelectAll()
			}
		}
   		return nil
   	}))
    obj.Set("onfocus",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		block := widget.Block()
		fmt.Printf("OnFocus:%p,%d (from %p,%d)\n",widget,block.Pos(),
			block.lastFocusOutWidget,block.lastFocusOutPos)
		if rownum != NOTINTABLE {
			buffer := block.Buffer()
			bufferPos,_ := block.ViewRow2BufferPos(rownum)
			if bufferPos != buffer.pos {
				buffer.Goto(bufferPos)
				widget.Block().Refresh()
				widget.SetFocus()
			}	
		}
		widget.Block().OnFocusToWidget(widget)
		widget.Block().RefreshCurrentRow()
		widget.SelectAll()
   		return nil
   	}))
	obj.Call("addEventListener","focusout",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		block := widget.Block()
		//fmt.Printf("OnFocusOut:%p,%d\n",widget,block.Pos())
		widget.WriteIfChanged(obj,rownum)
		block.RefreshCurrentRow()
		block.lastFocusOutWidget = widget
		block.lastFocusOutPos = block.Pos()
		// preventing focusout: widget.SetFocus()
   		return nil
   	}))

}



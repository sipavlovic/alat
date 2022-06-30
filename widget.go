
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


// ----------------------------------------------------------


func AttachFocusEvents(widget Widget, obj js.Value, rownum int) {
	obj.Set("onkeydown",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		keycode := event.Get("keyCode").Int()
		shiftkey := event.Get("shiftKey").Bool()
		ctrlkey := event.Get("ctrlKey").Bool()
		fmt.Println("OnKeyUp:",keycode,shiftkey,ctrlkey,rownum)
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
			}
		case 33: // PgUp
			event.Call("preventDefault")
			if !shiftkey && !ctrlkey && rownum != NOTINTABLE {
				widget.WriteIfChanged(obj,rownum)
				buffer := widget.Block().Buffer()
				table := widget.(*Column).parentWidget.(*Table)
				lastPos := buffer.pos
				if lastPos == table.viewBegin {
					buffer.Goto(lastPos-table.visibleRows)
				} else {
					buffer.Goto(table.viewBegin)
				}
				if lastPos != buffer.pos {
					widget.Block().Refresh()
					widget.SetFocus()
				}
			}
		case 34: // PgDn
			event.Call("preventDefault")
			if !shiftkey && !ctrlkey && rownum != NOTINTABLE {
				widget.WriteIfChanged(obj,rownum)
				buffer := widget.Block().Buffer()
				table := widget.(*Column).parentWidget.(*Table)
				lastPos := buffer.pos
				if lastPos == table.viewEnd {
					buffer.Goto(lastPos+table.visibleRows)
				} else {
					buffer.Goto(table.viewEnd)
				}
				if lastPos != buffer.pos {
					widget.Block().Refresh()
					widget.SetFocus()
				}
			}
		}
   		return nil
   	}))
    obj.Set("onfocus",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("OnFocus:",obj,rownum)
		if rownum != NOTINTABLE {
			table := widget.(*Column).parentWidget.(*Table)
			buffer := widget.Block().Buffer()
			bufferPos,_ := table.ViewRow2BufferPos(rownum)
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
		fmt.Println("OnFocusOut:",obj,rownum)
		widget.WriteIfChanged(obj,rownum)
		widget.Block().RefreshCurrentRow()
		// preventing focusout: widget.SetFocus()
   		return nil
   	}))

}


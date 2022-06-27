
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

// ----------------------------------------------------------


func AttachOnChangeEvent(widget Widget, obj js.Value, rownum int) {
	obj.Set("onchange",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if column,ok := widget.Block().widgetsToColumns[widget]; ok {
			value := obj.Get("value").String()
			if rownum < 0 {
				widget.Block().Buffer().Set(column,value)
			} else {
				table := widget.(*Column).parentWidget.(*Table)
				pos,_ := table.ViewRow2BufferPos(rownum)
				widget.Block().Buffer().SetAt(pos,column,value)
			}
			fmt.Println("OnChange",column,"to",value,"rownum",rownum)
		}	
   		return nil
   	}))
}


func AttachFocusEvents(widget Widget, obj js.Value, pos int) {
	obj.Set("onkeydown",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		keycode := event.Get("keyCode").Int()
		shiftkey := event.Get("shiftKey").Bool()
		ctrlkey := event.Get("ctrlKey").Bool()
		fmt.Println("OnKeyUp:",keycode,shiftkey,ctrlkey,pos)
		switch keycode {
		case 9: // Tab
			event.Call("preventDefault")
			if shiftkey && !ctrlkey {
				widget.Block().PrevWidget().SetFocus()
			} else if !shiftkey && !ctrlkey {
				widget.Block().NextWidget().SetFocus()
			}
		case 38: // Up
			event.Call("preventDefault")
			if !shiftkey && !ctrlkey {
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
			if !shiftkey && !ctrlkey && pos != NOTINTABLE {
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
			if !shiftkey && !ctrlkey && pos != NOTINTABLE {
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
		fmt.Println("OnFocus:",obj,pos)
		if pos != NOTINTABLE {
			table := widget.(*Column).parentWidget.(*Table)
			buffer := widget.Block().Buffer()
			bufferPos,_ := table.ViewRow2BufferPos(pos)
			if bufferPos != buffer.pos {
				buffer.Goto(bufferPos)
				widget.Block().Refresh()
				widget.SetFocus()
			}
		}
		widget.Block().OnFocusToWidget(widget)
   		return nil
   	}))
	obj.Call("addEventListener","focusout",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("OnFocusOut:",obj,pos)
   		return nil
   	}))

}


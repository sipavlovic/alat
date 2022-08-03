
package alat

import (
	"syscall/js"
)


func AttachFocusEvents(widget FocusableWidget, obj js.Value, rownum int) {
	
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
		case 36: // Home
			if !shiftkey && ctrlkey {
				event.Call("preventDefault")
				widget.WriteIfChanged(obj,rownum)
				block := widget.Block()
				buffer := block.Buffer()
				lastPos := buffer.pos
				buffer.Goto(0)
				if lastPos != buffer.pos {
					block.Refresh()
					widget.SetFocus()
				}
				widget.SelectAll()
			}
		case 35: // End
			if !shiftkey && ctrlkey {
				event.Call("preventDefault")
				widget.WriteIfChanged(obj,rownum)
				block := widget.Block()
				buffer := block.Buffer()
				lastPos := buffer.pos
				buffer.Goto(len(buffer.rows)-1)
				if lastPos != buffer.pos {
					block.Refresh()
					widget.SetFocus()
				}
				widget.SelectAll()
			}
		}
   		return nil
   	}))

	/*
	obj.Set("onmousedown",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Printf("OnClick:%p,%d\n",widget,rownum)
		return nil
	}))
	*/

    obj.Set("onfocus",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		block := widget.Block()
		//fmt.Printf("OnFocus:%p,%d (from %p,%d)\n",widget,block.Pos(),
		//	block.lastFocusOutWidget,block.lastFocusOutPos)
		if widget.IsInMultiRow() {
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



package alat

import (
	"syscall/js"
)

type Button struct {
	BaseFocusableWidget
	handler func(* Button)
}

func NewButton(block *Block, parentWidget ParentWidget, label string) *Button {
	var button Button
	button.BaseFocusableWidget.Init(block, &button, parentWidget, label)
	button.handler = nil
	return &button
}

func (w *Button) Draw() {
	button := NewNode(w.ParentHTMLObject(),"button")
	button.Set("textContent",w.Label())
	w.htmlObject = button
	w.BaseFocusableWidget.Draw()
	button.Set("onclick",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if w.handler != nil {
			w.handler(w)
		}
   		return nil
   	}))
	AttachFocusEvents(w,button,NOTINTABLE)
}

func (w *Button) DrawInMultiRow(parent js.Value, rownum int) js.Value {
	button := NewNode(parent,"button")
	button.Set("textContent",w.Label())
	button.Set("onclick",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if w.handler != nil {
			w.handler(w)
		}
   		return nil
   	}))
	AttachFocusEvents(w,button,rownum)
	return button
}

func (w *Button) SelectAll() {}

func (w *Button) RefreshCell(cell js.Value, rowState int, value string) {
	if rowState == ROWNUM_NOT_EXISTS {
		cell.Set("type","hidden")
	} else {
		cell.Set("type","")
	}
}

func (w *Button) SetHandler(hnd func(* Button)) {
	w.handler = hnd
}




package alat

import (
	"syscall/js"
)

type Button struct {
	BaseWidget
	text string
	handler func(* Button)
}

func NewButton(block *Block, parentWidget Widget, text string) *Button {
	var button Button
	button.BaseWidget.Init(block, &button, parentWidget)
	button.text = text
	block.AddToFocusList(&button)
	button.handler = nil
	return &button
}

func (w *Button) Draw() {
	button := NewNode(w.ParentHTMLObject(),"button")
	button.Set("textContent",w.text)
	w.htmlObject = button
	w.BaseWidget.Draw()
	button.Set("onclick",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if w.handler != nil {
			w.handler(w)
		}
   		return nil
   	}))
	AttachFocusEvents(w,button,NOTINTABLE)
}

func (w *Button) SetFocus() {
	w.HTMLObject().Call("focus")
}

func (w *Button) SetHandler(hnd func(* Button)) {
	w.handler = hnd
}


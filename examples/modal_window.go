
package main

import (
	"github.com/sipavlovic/alat"
	"syscall/js"
	"strconv"
)

type ModalWindow struct {
	alat.BaseParentWidget
	divWindow js.Value
	divHeader js.Value
	title string
	pos1 int
	pos2 int
	pos3 int
	pos4 int
	BackupDocumentOnMouseUp js.Value
	BackupDocumentOnMouseMove js.Value
}

func NewModalWindow(block *alat.Block, title string) *ModalWindow {
	var window ModalWindow
	window.BaseParentWidget.Init(block, &window, nil)
	window.title = title
	return &window
}


func (w *ModalWindow) Close() {
	alat.RemoveNode(w.divWindow)
}

func (w *ModalWindow) Draw() {
	w.divWindow = alat.NewNode(w.ParentHTMLObject(),"div")
	w.divWindow.Set("style","position:absolute;z-index:9;background-color:#fff;border:1px solid #888;")
	w.divWindow.Get("style").Set("top","100px")
	w.divWindow.Get("style").Set("left","100px")
	w.divHeader = alat.NewNode(w.divWindow,"div")
	w.divHeader.Set("style","padding:10px;cursor:move;z-index:10;background-color:#2196F3;color:#fff;")
	w.divHeader.Set("textContent",w.title)
	divContent := alat.NewNode(w.divWindow,"div")
	divContent.Set("style","padding:10px;")
	w.SetHTMLObject(divContent)
	w.defineDraggable()
	for _, child := range w.Children() {
		child.Draw()
	}
}

func (w *ModalWindow) defineDraggable() {
	
	dragMouseDown := func(event js.Value) {

		elementDrag := func(event js.Value) {
			event.Call("preventDefault")
			w.pos1 = w.pos3 - event.Get("clientX").Int()
			w.pos2 = w.pos4 - event.Get("clientY").Int()
			w.pos3 = event.Get("clientX").Int()
			w.pos4 = event.Get("clientY").Int()
			top := strconv.Itoa(w.divWindow.Get("offsetTop").Int()-w.pos2)
			left := strconv.Itoa(w.divWindow.Get("offsetLeft").Int()-w.pos1)
			w.divWindow.Get("style").Set("top",top+"px")
			w.divWindow.Get("style").Set("left",left+"px")
		}
		closeDragElement := func() {
			alat.HTMLDocument.Set("onmouseup",w.BackupDocumentOnMouseUp)
			alat.HTMLDocument.Set("onmousemove",w.BackupDocumentOnMouseMove)
		}

		event.Call("preventDefault")
		w.pos3 = event.Get("clientX").Int()
    	w.pos4 = event.Get("clientY").Int()
		w.BackupDocumentOnMouseUp = alat.HTMLDocument.Get("onmouseup")
		w.BackupDocumentOnMouseMove = alat.HTMLDocument.Get("onmousemove")
		alat.HTMLDocument.Set("onmouseup",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			closeDragElement()
			return nil
		}))
		alat.HTMLDocument.Set("onmousemove",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			elementDrag(args[0])
			return nil
		}))
	}

	w.divHeader.Set("onmousedown",js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		dragMouseDown(args[0])
		return nil
	}))
	bx := w.divWindow.Get("offsetLeft").String()
  	by := w.divWindow.Get("offsetTop").String()
	w.divWindow.Get("style").Set("top",by+"px")
	w.divWindow.Get("style").Set("left",bx+"px")
}




package alat

type Label struct {
	BaseWidget
	text string
}

func NewLabel(block *Block, parentWidget ParentWidget, text string) *Label {
	var label Label
	label.BaseWidget.Init(block, &label, parentWidget)
	label.text = text
	return &label
}

func (w *Label) Draw() {
	div := NewNode(w.ParentHTMLObject(),"div")
	div.Set("textContent",w.text)
	w.htmlObject = div
	w.BaseWidget.Draw()
}

func (w *Label) Remove() {
	RemoveNode(w.HTMLObject())
}



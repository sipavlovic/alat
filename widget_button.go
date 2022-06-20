package alat

type Button struct {
	BaseWidget
	text string
}

func NewButton(block *Block, parentWidget Widget, text string) *Button {
	var button Button
	button.BaseWidget.Init(block, &button, parentWidget)
	button.text = text
	return &button
}

func (w *Button) Draw() {
	button := NewNode(w.ParentHTMLObject(),"button")
	button.Set("textContent",w.text)
	w.htmlObject = button
	w.BaseWidget.Draw()
}
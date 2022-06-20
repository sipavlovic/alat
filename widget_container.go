package alat

type Container struct {
	BaseWidget
}

func NewContainer(block *Block, parentWidget Widget) *Container {
	var container Container
	container.BaseWidget.Init(block, &container, parentWidget)
	return &container
}

func (w *Container) Draw() {
	div := NewNode(w.ParentHTMLObject(),"div")
	w.htmlObject = div
	//w.BaseWidget.Draw()
	for _, child := range w.children {
		child.Draw()
	}
}

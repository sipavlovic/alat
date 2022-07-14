package alat

type Container struct {
	BaseParentWidget
}

func NewContainer(block *Block, parentWidget ParentWidget) *Container {
	var container Container
	container.BaseParentWidget.Init(block, &container, parentWidget)
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

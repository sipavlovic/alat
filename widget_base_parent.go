
package alat

type ParentWidget interface {
	Widget
	AddChild(Widget)
	Children() []Widget 
}

type BaseParentWidget struct {
	BaseWidget
	children []Widget
}

func (w *BaseParentWidget) Init(block *Block, widget Widget, parentWidget ParentWidget) {
	w.BaseWidget.Init(block, widget, parentWidget)
}

func (w *BaseParentWidget) IsParent() bool {
	return true
}

func (w *BaseParentWidget) AddChild(child Widget) {
	w.children = append(w.children,child)
}

func (w *BaseParentWidget) Children() []Widget {
	return w.children
}

func (w *BaseParentWidget) Draw() {
	for _, child := range w.children {
		child.Draw()
	}
}

func (w *BaseParentWidget) Refresh() {
	for _, child := range w.children {
		child.Refresh()
	}
}

func (w *BaseParentWidget) RefreshCurrentRow() {
	for _, child := range w.children {
		child.RefreshCurrentRow()
	}
}





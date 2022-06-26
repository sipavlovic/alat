
package alat

type Edit struct {
	BaseWidget
}

func NewEdit(block *Block, parentWidget Widget) *Edit {
	var edit Edit
	edit.BaseWidget.Init(block, &edit, parentWidget)
	return &edit
}

func (w *Edit) Draw() {
	input := NewNode(w.ParentHTMLObject(),"input")
	w.htmlObject = input
	w.BaseWidget.Draw()
}

func (w *Edit) Refresh() {
	result := ""
	if column,ok := w.Block().widgetsToColumns[w]; ok {
		result,_ = w.Block().Buffer().Get(column)
	}
	w.HTMLObject().Set("value",result)
	w.BaseWidget.Refresh()
}


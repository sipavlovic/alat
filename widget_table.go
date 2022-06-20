package alat

import (
	"syscall/js"
)

type Table struct {
	BaseWidget
	visibleRows int
	tableObj js.Value
	columns []*Column
	viewBegin int
	viewEnd int
}

func NewTable(block *Block, parentWidget Widget, visibleRows int) *Table {
	var table Table
	table.BaseWidget.Init(block, &table, parentWidget)
	table.visibleRows = visibleRows
	return &table
}

func (w *Table) AddColumn(label string) *Column {
	var col Column
	col.BaseWidget.Init(w.block, &col, w)
	col.label = label
	w.columns = append(w.columns,&col)
	col.index = len(w.columns)-1
	return &col
}

func (w *Table) Draw() {
	div := NewNode(w.ParentHTMLObject(),"div")
	w.htmlObject = div
	div.Set("hidden",false)
	div.Set("style","width: 600px; overflow:auto")
	w.tableObj = NewNode(div,"table")
	w.DrawContent()
}

func (w *Table) DrawContent() {
	ClearNode(w.tableObj)
	//obj.viewBegin, obj.viewEnd = obj.block.Buffer.CalcView(obj.viewBegin, obj.viewEnd, obj.viewRows)
	tr := NewNode(w.tableObj,"tr")
	for _,col := range w.columns {
		th := NewNode(tr,"th")
		th.Set("textContent",col.label)
	}
	for rownum:=0;rownum<w.visibleRows;rownum++ {
		w.DrawRow(rownum)
	}
}


func (w *Table) DrawRow(rownum int) {
	rowtr := NewNode(w.tableObj,"tr")
	for range w.columns {
		rowtd := NewNode(rowtr,"td")
		NewNode(rowtd,"input")
	}
}








package alat

import (
	"syscall/js"
	"strconv"
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
	w.Block().AddToFocusList(&col)
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


func (w *Table) Refresh() {
	w.DrawContent()
}


func (w *Table) DrawContent() {
	ClearNode(w.tableObj)
	w.viewBegin, w.viewEnd = w.Block().Buffer().CalcView(w.viewBegin, w.viewEnd, w.visibleRows)
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
	bufferRow := w.viewBegin+rownum
	rowtr := NewNode(w.tableObj,"tr")
	if bufferRow>=0 && bufferRow<=w.viewEnd {
		for _,columnWidget := range w.columns {
			rowtd := NewNode(rowtr,"td")
			input := NewNode(rowtd,"input")
			column := w.Block().widgetsToColumns[columnWidget]
			value,_ := w.Block().Buffer().GetAt(bufferRow,column)
			input.Set("value", value)
			AttachOnChangeEvent(columnWidget,input,rownum)
			AttachFocusEvents(columnWidget,input,rownum)
		}
	} else {
		for range w.columns {
			rowtd := NewNode(rowtr,"td")
			rowtd.Set("style","background-color: #CCC")
		}
	}
}


func (w *Table) Cell(colnum int, rownum int) js.Value {
	tr := w.tableObj.Get("children").Get(strconv.Itoa(rownum+1))
	td := tr.Get("children").Get(strconv.Itoa(colnum))
	return td.Get("children").Get("0")
}

func (w *Table) ViewRow2BufferPos(viewRow int) (int, bool) {
	if viewRow>=0 && viewRow<w.visibleRows {
		buffer := w.Block().Buffer()
		bufferPos := viewRow+w.viewBegin
		if bufferPos>=0 && bufferPos<len(buffer.rows) {
			return bufferPos,true
		}
	}
	return 0,false
}

func (w *Table) BufferPos2ViewRow(bufferPos int) (int, bool) {
	buffer := w.Block().Buffer()
	if bufferPos>=0 && bufferPos<len(buffer.rows) {
		viewRow := bufferPos-w.viewBegin
		if viewRow>=0 && viewRow<w.visibleRows {
			return viewRow,true
		}
	}
	return 0,false
}






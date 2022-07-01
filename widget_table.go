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
	w.viewBegin, w.viewEnd = w.Block().Buffer().CalcView(w.viewBegin, w.viewEnd, w.visibleRows)
	for rownum:=0;rownum<w.visibleRows;rownum++ {
		w.RefreshRownum(rownum)
	}
}

func (w *Table) RefreshCurrentRow() {
	pos := w.Block().Buffer().pos
	rownum,_ := w.BufferPos2ViewRow(pos)
	w.RefreshRownum(rownum)
}


func (w *Table) RefreshRownum(rownum int) {
	for _,widgetColumn := range w.columns {
		w.RefreshColumnAtRownum(widgetColumn,rownum)
	}
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
	rowtr := NewNode(w.tableObj,"tr")
	for _,columnWidget := range w.columns {
		rowtd := NewNode(rowtr,"td")
		input := NewNode(rowtd,"input")
		AttachFocusEvents(columnWidget,input,rownum)
	}	
}


func (w *Table) Cell(colnum int, rownum int) js.Value {
	tr := w.tableObj.Get("children").Get(strconv.Itoa(rownum+1))
	td := tr.Get("children").Get(strconv.Itoa(colnum))
	return td.Get("children").Get("0")
}


func (w *Table) RefreshColumnAtRownum(widgetColumn *Column, rownum int) {
	bufferRow := w.viewBegin+rownum
	pos := w.Block().Buffer().pos
	input := w.Cell(widgetColumn.index,rownum)
	td := input.Get("parentElement")
	result := ""
	if bufferRow>=0 && bufferRow<=w.viewEnd {
		if column,ok := w.Block().widgetsToColumns[widgetColumn]; ok {
			result,_ = w.Block().Buffer().GetAt(bufferRow,column)
		}
		if bufferRow==pos {
			td.Set("style","background-color: #BCE5FD")
			input.Set("style","background-color: #BCE5FD")
			input.Set("type","")
		} else {
			td.Set("style","background-color: #FFFFFF")
			input.Set("style","background-color: #FFFFFF")
			input.Set("type","")
		}
	} else {
		td.Set("style","background-color: #CCC")
		input.Set("style","background-color: #CCC")
		input.Set("type","hidden")
	}
	input.Set("value",result)
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






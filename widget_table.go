package alat

import (
	"syscall/js"
	"strconv"
)

type Table struct {
	BaseWidget
	tableObj js.Value
	columns []*Column
}

func NewTable(block *Block, parentWidget Widget) *Table {
	var table Table
	table.BaseWidget.Init(block, &table, parentWidget)
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
	block := w.Block()
	block.viewBegin, block.viewEnd = block.Buffer().CalcView(block.viewBegin, block.viewEnd, block.visibleRows)
	for rownum:=0;rownum<block.visibleRows;rownum++ {
		w.RefreshRownum(rownum)
	}
}

func (w *Table) RefreshCurrentRow() {
	pos := w.Block().Buffer().pos
	rownum,_ := w.Block().BufferPos2ViewRow(pos)
	w.RefreshRownum(rownum)
}


func (w *Table) RefreshRownum(rownum int) {
	for _,widgetColumn := range w.columns {
		w.RefreshColumnAtRownum(widgetColumn,rownum)
	}
}


func (w *Table) DrawContent() {
	ClearNode(w.tableObj)
	block := w.Block()
	block.viewBegin, block.viewEnd = block.Buffer().CalcView(block.viewBegin, block.viewEnd, block.visibleRows)
	tr := NewNode(w.tableObj,"tr")
	for _,col := range w.columns {
		th := NewNode(tr,"th")
		th.Set("textContent",col.label)
	}
	for rownum:=0;rownum<block.visibleRows;rownum++ {
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
	block := w.Block() 
	bufferRow := block.viewBegin+rownum
	pos := block.Buffer().pos
	input := w.Cell(widgetColumn.index,rownum)
	td := input.Get("parentElement")
	result := ""
	if bufferRow>=0 && bufferRow<=block.viewEnd {
		if column,ok := block.widgetsToColumns[widgetColumn]; ok {
			result,_ = block.Buffer().GetAt(bufferRow,column)
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








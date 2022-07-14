package alat

import (
	"syscall/js"
	"strconv"
)

type Table struct {
	BaseMultiRowWidget
	tableObj js.Value
}

func NewTable(block *Block, parentWidget ParentWidget) *Table {
	var table Table
	table.BaseMultiRowWidget.Init(block, &table, parentWidget)
	return &table
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
	block := w.Block()
	block.viewBegin, block.viewEnd = block.Buffer().CalcView(block.viewBegin, block.viewEnd, block.visibleRows)
	tr := NewNode(w.tableObj,"tr")
	for _,col := range w.Children() {
		th := NewNode(tr,"th")
		th.Set("textContent",col.(FocusableWidget).Label())
	}
	for rownum:=0;rownum<block.visibleRows;rownum++ {
		w.DrawRow(rownum)
	}
}

func (w *Table) DrawRow(rownum int) {
	rowtr := NewNode(w.tableObj,"tr")
	for _,columnWidget := range w.Children() {
		rowtd := NewNode(rowtr,"td")
		columnWidget.(FocusableWidget).DrawInMultiRow(rowtd,rownum)
	}	
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
	for _,widgetColumn := range w.Children() {
		w.RefreshColumnAtRownum(widgetColumn.(FocusableWidget),rownum)
	}
}

func (w *Table) RefreshColumnAtRownum(widgetColumn FocusableWidget, rownum int) {
	block := w.Block() 
	bufferRow := block.viewBegin+rownum
	pos := block.Buffer().pos
	td := w.CellElement(widgetColumn.Index(),rownum)
	columnCell := widgetColumn.ColumnCell(td)
	result := ""
	if bufferRow>=0 && bufferRow<=block.viewEnd {
		if column,ok := block.widgetsToColumns[widgetColumn]; ok {
			result,_ = block.Buffer().GetAt(bufferRow,column)
		}
		if bufferRow==pos {
			td.Set("style","background-color: #BCE5FD")
			widgetColumn.RefreshCell(columnCell,ROWNUM_CURRENT,result)
		} else {
			td.Set("style","background-color: #FFFFFF")
			widgetColumn.RefreshCell(columnCell,ROWNUM_NOT_CURRENT,result)
		}
	} else {
		td.Set("style","background-color: #CCC")
		widgetColumn.RefreshCell(columnCell,ROWNUM_NOT_EXISTS,result)
	}
}

func (w *Table) CellElement(colnum int, rownum int) js.Value {
	tr := w.tableObj.Get("children").Get(strconv.Itoa(rownum+1))
	td := tr.Get("children").Get(strconv.Itoa(colnum))
	return td
}




package alat

import (
	"fmt"
)


type Buffer struct {
	pos int
	rows []Row
}	


func NewBuffer() *Buffer {
	b := Buffer{}
	b.pos = NOROW
	b.rows = make([]Row,0)
	return &b
}


func (b *Buffer) InsertRow() int { 
	b.rows = append(b.rows, MakeRow())
	b.pos = b.Goto(len(b.rows)-1)
	return b.pos
}


func (b *Buffer) Get(column string) (string,error) {
	return b.GetAt(b.pos,column)
}


func (b* Buffer) Set(column string, value string) error {
	return b.SetAt(b.pos,column,value)
}


func (b *Buffer) GetAt(pos int, column string) (string,error) {
	if len(b.rows)>0 && pos>=0 && pos<len(b.rows) {
		if value,ok := b.rows[pos].Values[column]; ok {
			return value,nil
		}
	}
	return "",fmt.Errorf("GetAt(%d,%s): no value",pos,column)
}


func (b* Buffer) SetAt(pos int, column string, value string) error {
	if len(b.rows)>0 && pos>=0 && pos<len(b.rows) {
		b.rows[pos].Values[column] = value
		return nil
	}
	return fmt.Errorf("SetAt(%d,%s,%s): error setting value",pos,column)
}


func (b* Buffer) Goto(newPos int) int {
	b.pos = newPos
	if len(b.rows)>0 {
		if b.pos < 0 {
			b.pos = 0
		} else if b.pos >= len(b.rows) {
			b.pos = len(b.rows)-1
		}
	} else {
		b.pos = NOROW
	}
	return b.pos
}


func (b *Buffer) CalcView(begin int, end int, rows int) (int,int) {
	// return begin,end
	newBegin := begin
	newEnd := end	
	// If there is no rows in buffer
	if len(b.rows)==0 {
		return NOROW,NOROW
	}
	// If num of rows in buffers are less or equal to rows in view
	if len(b.rows)<=rows {
		return 0,len(b.rows)-1
	}
	// If begin is greater than pos
	if newBegin>b.pos {
		newBegin = b.pos
		newEnd = newBegin+rows-1
	}
	// if end is lesser than pos
	if newEnd<b.pos {
		newEnd = b.pos
		newBegin = newEnd-(rows-1)
	}
	// if begin < 0
	if newBegin<0 {
		newBegin = 0
		newEnd = newBegin+rows-1
	}
	// if end >= len rows
	if newEnd>=len(b.rows) {
		newEnd = len(b.rows)-1
		newBegin = newEnd-(rows-1)
	}
	return newBegin,newEnd
}



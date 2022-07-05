
/*
LOCAL MODULE:
go mod edit -replace github.com/sipavlovic/alat=../../alat
*/

package main

import (
	"github.com/sipavlovic/alat"
	"syscall/js"
	"fmt"
	"strconv"
)

const (
	version = "2022-07-05"
)

func MakeCSS() {
	head := js.Global().Get("document").Get("head")
	css := alat.NewNode(head,"style")
	css.Set("textContent",`
	table, td, th { 
		border: 1px solid #CCC; 
		border-collapse: collapse; 
		}	

	td, th { 
		background-color: #FFF; 
		height: 20px; 
		text-align: center; 
		vertical-align: middle;
		}		

	th {
		background-color: #333;
		color: #FFF;
		border-color: #555;
		}

	td input {
		border: 0px solid #CCC; 
	}	

	`)
}


func ToInt(value string) (int,error) {
	return strconv.Atoi(value)
}
func FromInt(value int) string {
	return strconv.Itoa(value)
}


type NewLineWidget struct {
	alat.BaseWidget
}
func NewLine(block *alat.Block, parentWidget alat.Widget) *NewLineWidget {
	var w NewLineWidget
	w.BaseWidget.Init(block, &w, parentWidget)
	return &w
}
func (w *NewLineWidget) Draw() {
	elem := alat.NewNode(w.ParentHTMLObject(),"BR")
	w.BaseWidget.SetHTMLObject(elem)
	w.BaseWidget.Draw()
}



func main() {

	MakeCSS()
	body := js.Global().Get("document").Get("body")

	block := alat.NewBlock(body,10)
	container := alat.NewContainer(block,nil) 
	alat.NewLabel(block,container,"Alat Example version: "+version)
	NewLine(block,container)
	alat.NewLabel(block,container,"Enter username:")
	w_usr := alat.NewEdit(block,container)
	alat.NewLabel(block,container,"Enter password:")
	w_pwd := alat.NewEdit(block,container)
	alat.NewLabel(block,container,"Password (copy):")
	w_pwd2 := alat.NewEdit(block,container)
	NewLine(block,container)
	NewLine(block,container)
	alat.NewLabel(block,container,"Table:")
	table := alat.NewTable(block,container)
	col_one := table.AddColumn("One")
	col_one2 := table.AddColumn("One (copy)")
	col_two := table.AddColumn("Two")
	col_three := table.AddColumn("Three")
	col_four := table.AddColumn("Four")
	col_five := table.AddColumn("Five")
	col_six := table.AddColumn("Six")
	NewLine(block,container)
	butt := alat.NewButton(block,container,"Click on me!")
	butt.SetHandler(func(w *alat.Button) {
		js.Global().Call("alert","Clicked!")
	})

	block.Connect(w_usr,"USERNAME")
	block.Connect(w_pwd,"PASSWORD")
	block.Connect(w_pwd2,"PASSWORD")
	block.Connect(col_one,"ONE")
	block.Connect(col_one2,"ONE")
	block.Connect(col_two,"TWO")
	block.Connect(col_three,"THREE")
	block.Connect(col_four,"FOUR")
	block.Connect(col_five,"FIVE")
	block.Connect(col_six,"SIX")
	
	buff := block.Buffer()
	for i:=1;i<=50;i++ {
		buff.InsertRow()
		buff.Set("ID",FromInt(i))
		buff.Set("USERNAME","User-"+FromInt(i))
		buff.Set("PASSWORD","Pwd-"+FromInt(i))
		buff.Set("ONE","One-"+FromInt(i))
		buff.Set("TWO","Two-"+FromInt(i))
		buff.Set("THREE","Three-"+FromInt(i))
		buff.Set("FOUR","Four-"+FromInt(i))
		buff.Set("FIVE","Five-"+FromInt(i))
		buff.Set("SIX","Six-"+FromInt(i))
	}

	block.Draw()
	block.Refresh()
	
	fmt.Println("End defining.")

	<-make(chan struct{})
}
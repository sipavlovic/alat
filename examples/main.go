
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



func main() {

	MakeCSS()
	body := js.Global().Get("document").Get("body")

	block := alat.NewBlock(body)
	container := alat.NewContainer(block,nil) 
	alat.NewLabel(block,container,"Enter username:")
	w_usr := alat.NewEdit(block,container)
	alat.NewLabel(block,container,"Enter password:")
	w_pwd := alat.NewEdit(block,container)
	alat.NewLabel(block,container,"Table:")
	table := alat.NewTable(block,container,6)
	col_one := table.AddColumn("One")
	col_two := table.AddColumn("Two")
	col_three := table.AddColumn("Three")
	col_four := table.AddColumn("Four")
	col_five := table.AddColumn("Five")
	col_six := table.AddColumn("Six")
	alat.NewButton(block,container,"Click on me!")

	block.Connect(w_usr,"USERNAME")
	block.Connect(w_pwd,"PASSWORD")
	block.Connect(col_one,"ONE")
	block.Connect(col_two,"TWO")
	block.Connect(col_three,"THREE")
	block.Connect(col_four,"FOUR")
	block.Connect(col_five,"FIVE")
	block.Connect(col_six,"SIX")
	
	buff := block.Buffer()
	for i:=1;i<=10;i++ {
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
	fmt.Println("Buffer:",buff)

	block.Draw()
	fmt.Println("End defining.")

	<-make(chan struct{})
}
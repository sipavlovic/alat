
/*
LOCAL MODULE:
go mod edit -replace github.com/sipavlovic/alat=../../alat
*/

package main

import (
	"github.com/sipavlovic/alat"
	"syscall/js"
	"fmt"
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


func main() {

	MakeCSS()
	body := js.Global().Get("document").Get("body")

	block := alat.NewBlock(body)
	container := alat.NewContainer(block,nil) 
	alat.NewLabel(block,container,"Enter username:")
	alat.NewEdit(block,container)
	alat.NewLabel(block,container,"Enter password:")
	alat.NewEdit(block,container)
	alat.NewLabel(block,container,"Table:")
	table := alat.NewTable(block,container,6)
	table.AddColumn("One")
	table.AddColumn("Two")
	table.AddColumn("Three")
	table.AddColumn("Four")
	table.AddColumn("Five")
	table.AddColumn("Six")

	alat.NewButton(block,container,"Click on me!")


	block.Draw()
	fmt.Println("End defining.")

	<-make(chan struct{})
}
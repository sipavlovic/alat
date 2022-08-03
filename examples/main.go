
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
	version = "2022-08-03"
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
func NewLine(block *alat.Block, parentWidget alat.ParentWidget) *NewLineWidget {
	var w NewLineWidget
	w.BaseWidget.Init(block, &w, parentWidget)
	return &w
}
func (w *NewLineWidget) Draw() {
	elem := alat.NewNode(w.ParentHTMLObject(),"BR")
	w.BaseWidget.SetHTMLObject(elem)
	w.BaseWidget.Draw()
}

func (w *NewLineWidget) Remove() {
	alat.RemoveNode(w.HTMLObject())
}












// ====================================================================================

func CallMainBlock() string {

	block := alat.NewBlock(alat.HTMLBody,10)
	container := alat.NewContainer(block,nil) 
	alat.NewLabel(block,container,"Alat Example version: "+version)
	NewLine(block,container)
	w_usr := alat.NewEdit(block,container,"Enter username:")
	w_pwd := alat.NewEdit(block,container,"Enter password:")
	w_pwd2 := alat.NewEdit(block,container,"Password (copy):")
	NewLine(block,container)
	NewLine(block,container)

	alat.NewLabel(block,container,"Table:")
	table := alat.NewTable(block,container)
	col_one := alat.NewEdit(block,table,"One")
	col_one2 := alat.NewEdit(block,table,"One (copy)")
	col_two := alat.NewEdit(block,table,"Two")
	col_three := alat.NewEdit(block,table,"Three")
	col_four := alat.NewEdit(block,table,"Four")
	col_five := alat.NewEdit(block,table,"Five")
	col_six := alat.NewEdit(block,table,"Six")
	col_seven := alat.NewButton(block,table,"Click!")
	col_seven.SetHandler(func(w *alat.Button) {
		js.Global().Call("alert",fmt.Sprintf("Row pos: %d",block.Pos()))
	})
	NewLine(block,container)
	
	butt := alat.NewButton(block,container,"Open modal window!")
	butt.SetHandler(func(w *alat.Button) {
		msg := CallSubBlock()
		fmt.Println("Msg from Sub:",msg)
		fmt.Println("EXIT BUTTON CLICK ON MAIN BLOCK!!!")
	})

	alat.NewBlockField(block,"USERNAME")
	alat.NewRowField(block,"PASSWORD")

	block.Connect(w_usr,"USERNAME")
	block.Connect(w_pwd,"PASSWORD")
	block.Connect(w_pwd2,"PASSWORD")

	alat.NewRowField(block,"ONE")
	alat.NewRowField(block,"TWO")
	alat.NewRowField(block,"THREE")
	alat.NewRowField(block,"FOUR")
	alat.NewRowField(block,"FIVE")
	alat.NewRowField(block,"SIX")
	
	block.Connect(col_one,"ONE")
	block.Connect(col_one2,"ONE")
	block.Connect(col_two,"TWO")
	block.Connect(col_three,"THREE")
	block.Connect(col_four,"FOUR")
	block.Connect(col_five,"FIVE")
	block.Connect(col_six,"SIX")
		
	buff := block.Buffer()
	for _,field := range buff.Fields() {
		if !field.IsRowField() {
			buff.Set(field.Name(),field.Name()+"-Block")
		}
	}
	for i:=1;i<=50;i++ {
		buff.InsertRow()
		for _,field := range buff.Fields() {
			if field.IsRowField() {
				buff.Set(field.Name(),field.Name()+"-"+FromInt(i))
			}
		}
	}

	block.Draw()
	block.Refresh()

	fmt.Println("End defining main block.")

	return block.Wait()
}









// ==================================================================================

func CallSubBlock() string {

	block := alat.NewBlock(alat.HTMLBody,5)
	modal := alat.NewModalWindow(block,"Modal Window")

	alat.NewLabel(block,modal,"Some important messages...")
	NewLine(block,modal)
	table := alat.NewTable(block,modal)
	col_one := alat.NewEdit(block,table,"One")
	col_two := alat.NewEdit(block,table,"Two")
	col_three := alat.NewEdit(block,table,"Three")
	col_four := alat.NewEdit(block,table,"Four")
	col_five := alat.NewEdit(block,table,"Five")
	col_six := alat.NewEdit(block,table,"Six")
	NewLine(block,modal)

	butt := alat.NewButton(block,modal,"Close")
	butt.SetHandler(func(w *alat.Button) {
		block.Close("Closing sub block (from button), bro!")
	})
	modal.SetCloseHandler(func(w *alat.ModalWindow) {
		block.Close("Closing sub block (from X on window), bro!")
	})

	butt2 := alat.NewButton(block,modal,"New window")
	butt2.SetHandler(func(w *alat.Button) {
		msg := CallSubBlock()
		fmt.Println("Msg from Sub:",msg)
	})

	alat.NewRowField(block,"ONE")
	alat.NewRowField(block,"TWO")
	alat.NewRowField(block,"THREE")
	alat.NewRowField(block,"FOUR")
	alat.NewRowField(block,"FIVE")
	alat.NewRowField(block,"SIX")

	block.Connect(col_one,"ONE")
	block.Connect(col_two,"TWO")
	block.Connect(col_three,"THREE")
	block.Connect(col_four,"FOUR")
	block.Connect(col_five,"FIVE")
	block.Connect(col_six,"SIX")

	buff := block.Buffer()
	for i:=1;i<=100;i++ {
		buff.InsertRow()
		for _,field := range buff.Fields() {
			buff.Set(field.Name(),field.Name()+"-"+FromInt(i))
		}
	}

	block.Draw()
	block.Refresh()

	fmt.Println("End defining sub block.")

	return block.Wait()

}





// -----------------------------------------------

func main() {

	MakeCSS()

	msg := CallMainBlock()
	fmt.Println("Msg from Main:",msg)

	fmt.Println("THE END")

	<-make(chan struct{})
}
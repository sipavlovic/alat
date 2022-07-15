
package alat

import (
	"syscall/js"
)


func NewNode(parent js.Value, nodeType string) js.Value {
	node := js.Global().Get("document").Call("createElement",nodeType)
	parent.Call("appendChild", node)
	return node
}


func ClearNode(node js.Value) {
	node.Set("textContent","")
}

func RemoveNode(node js.Value) {
	node.Get("parentNode").Call("removeChild",node)
}



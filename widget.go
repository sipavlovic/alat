
package alat

import (
	"syscall/js"
)

func Test() {
	js.Global().Call("alert","Hello World!")
}

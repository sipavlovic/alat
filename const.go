
package alat

import (
	"syscall/js"
)

const (
	NOROW = -1
	NOTINTABLE = -2
	ROWNUM_NOT_EXISTS = 1
	ROWNUM_NOT_CURRENT = 2
	ROWNUM_CURRENT = 3
)

var (
	HTMLWindow = js.Global()
	HTMLDocument = HTMLWindow.Get("document")
	HTMLBody = HTMLDocument.Get("body")
)

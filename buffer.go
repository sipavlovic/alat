
package alat

type Row struct {
	Values map[string]string
}

type Buffer struct {
	Pos int
	Rows []Row
}	

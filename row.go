package alat

type Row struct {
	Values map[string]string
}

func MakeRow() Row {
	row := Row{}
	row.Values = make(map[string]string)
	return row 
}
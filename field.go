package alat

type Field struct {
	name string
	isRowField bool
	onChangeHandler func(*Field) bool
	onExitHandler func(*Field) bool
	onActionHandler func(*Field) bool
}

func NewBlockField(block *Block, name string) *Field { 
	var f Field
	f.name = name
	f.isRowField = false
	block.buffer.fields[name] = &f
	return &f
}

func NewRowField(block *Block, name string) *Field { 
	var f Field
	f.name = name
	f.isRowField = true
	block.buffer.fields[name] = &f
	return &f
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) IsRowField() bool {
	return f.isRowField
}
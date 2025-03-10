package components

import "github.com/a-h/templ"

type Table struct {
	Title    string
	Subtitle string
	Headers  []string
	Rows     [][]templ.Component
}

func (data Table) Component() templ.Component {
	return TableComponent(data)
}

type TableTextCell struct {
	Text string
}

func (data TableTextCell) Component() templ.Component {
	return tableCellComponent(false)
}

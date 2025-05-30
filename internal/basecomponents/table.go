package basecomponents

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

type TableLinkCell struct {
	Text string
	URL  string
	templ.Attributes
}

func (data TableTextCell) Component() templ.Component {
	return tableTextCellComponent(data)
}

func (data TableLinkCell) Component() templ.Component {
	return tableLinkCellComponent(data)
}

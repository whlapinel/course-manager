package domain

import (
	"fmt"
	"strconv"
)

type StandardSet struct {
	ID         int
	CourseName string
	Standards  []Standard
}
type Standard struct {
	Objective
	Children []Objective // objectives, etc
	StdSet   StandardSet // all standards should be associated with a set id
}

type Objective struct {
	ID          int
	ParentID    int    // negative if no parent
	Number      int    // number should not include parent number
	Name        string // official, from the state
	Description string // unofficial, teacher or PLC created
}

func (std Standard) Designation(standard, parent Standard) string {
	if std.ParentID == 0 {
		return strconv.Itoa(standard.Number)
	}
	return fmt.Sprintf("%d.%d", standard.Number, parent.Number)
}

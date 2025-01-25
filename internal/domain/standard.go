package domain

import "fmt"

type StandardSet struct {
	ID         int
	CourseName string
}
type Standard struct {
	StdSet      StandardSet // all standards should be associated with a set id
	ParentID    int         // negative if no parent
	Number      int         // number should not include parent number
	Name        string      // official, from the state
	Description string      // unofficial, teacher or PLC created
	Children    []Standard  // objectives, etc
}

func (std Standard) Designation(standard, parent Standard) string {
	return fmt.Sprintf("%d.%d", standard.Number, parent.Number)
}

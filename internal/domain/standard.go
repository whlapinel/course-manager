package domain

import "fmt"

type Standard struct {
	CourseID    int        // all standards should be associated with a course id
	ParentID    int        // negative if no parent
	Number      int        // number should not include parent number
	Name        string     // official, from the state
	Description string     // unofficial, teacher or PLC created
	Children    []Standard // objectives, etc
}

func (std Standard) Designation(standard, parent Standard) string {
	return fmt.Sprintf("%d.%d", standard.Number, parent.Number)
}

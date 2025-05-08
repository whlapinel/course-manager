package standard

import "fmt"

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
	ParentNum   int    // 0 if no parent
	Number      int    // number should not include parent number
	Name        string // official, from the state
	Description string // unofficial, teacher or PLC created
}

func (o Objective) Designation() string {
	return fmt.Sprintf("%d.%d", o.ParentNum, o.Number)
}

// implements node sorter
func (obj Objective) GetNumber() int {
	return obj.Number
}

type Objectives []Objective
type Standards []Standard

package standard

import "fmt"

type StandardSet struct {
	ID         int        `json:"id"`
	CourseName string     `json:"courseName"`
	Standards  []Standard `json:"standards"`
}

type Standard struct {
	Objective   `json:"objective"`
	Objectives  []Objective          `json:"objectives"` // objectives, etc
	StandardSet `json:"standardSet"` // all standards should be associated with a set id
}

type Objective struct {
	ID          int    `json:"id"`
	ParentID    int    `json:"parentID"`    // negative if no parent
	ParentNum   int    `json:"parentNum"`   // 0 if no parent
	Number      int    `json:"number"`      // number should not include parent number
	Name        string `json:"name"`        // official, from the state
	Description string `json:"description"` // unofficial, teacher or PLC created
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

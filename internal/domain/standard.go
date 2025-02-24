package domain

import (
	"fmt"
	"slices"
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
	ParentNum   int    // 0 if no parent
	Number      int    // number should not include parent number
	Name        string // official, from the state
	Description string // unofficial, teacher or PLC created
}

// implements node sorter
func (obj Objective) GetNumber() int {
	return obj.Number
}

type Objectives []Objective
type Standards []Standard

// implements node sorter
func (std Standard) GetNumber() int {
	return std.Number
}
func (std Standard) Designation() string {
	if std.ParentNum == 0 {
		return fmt.Sprintf("%d", std.Number)
	}
	return fmt.Sprintf("%d.%d", std.ParentNum, std.Number)
}

func (stds Objectives) Sort() {
	slices.SortFunc(stds, func(a, b Objective) int {
		if a.ParentNum < b.ParentNum {
			return -1
		}
		if a.ParentNum > b.ParentNum {
			return 1
		}
		if a.Number < b.Number {
			return -1
		}
		if a.Number > b.Number {
			return 1
		} else {
			return 0
		}
	})
}

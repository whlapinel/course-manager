package unit

import (
	"fmt"
	"gh_static_portfolio/internal/ports"
	"slices"
)

type Unit struct {
	ports.BaseNode[int, int]
}

func (u Unit) Designation() string {
	if u.Number < 0 {
		return "N/A"
	}
	return fmt.Sprintf("Unit %d", u.Number)
}

type Units []Unit

func (u Units) Sort() {
	slices.SortFunc[Units](u, func(a, b Unit) int {
		if a.Sequence > b.Sequence {
			return 1
		} else if b.Sequence > a.Sequence {
			return -1
		} else if a.Number > b.Number {
			return 1
		} else if b.Number > a.Number {
			return -1
		}
		return 0
	})

}

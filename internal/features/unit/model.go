package unit

import (
	"fmt"
	"gh_static_portfolio/internal/ports"
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

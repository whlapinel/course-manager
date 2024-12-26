package main

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

func ErrorMsg(err error) *widget.Label {
	return widget.NewLabel(fmt.Sprintf("error: %s", err))
}

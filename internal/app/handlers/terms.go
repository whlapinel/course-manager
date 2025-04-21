package handlers

import (
	"gh_static_portfolio/internal/features/course"
	"gh_static_portfolio/internal/features/term"
)

type TermHandler struct {
	termService term.Service
	courseService course.Service
}

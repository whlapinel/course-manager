package handlers

import (
	"gh_static_portfolio/internal/app/dto"
	"gh_static_portfolio/internal/app/services"
	"gh_static_portfolio/internal/shared/web"
)

type assessmentHandler struct {
	service     *services.AssessmentService
	nodeService *services.NodeService
	fileService *services.FileService
	markdown    *services.MarkdownService
	reverse     web.Reverse
	*baseHandler[dto.Course, int, int]
}

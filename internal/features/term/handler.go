package term

type Handler struct {
	reverse Reverse
	service Service
}

type Reverse func(name string, params ...any) string

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

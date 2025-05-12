package ports

type SlideRenderer interface {
	GetSlides(nodes ...Node) ([]byte, error)
	NewGetSlides(url string) ([]byte, error)
}

package ports

type SlideRenderer interface {
	GetSlides(url string) ([]byte, error)
}

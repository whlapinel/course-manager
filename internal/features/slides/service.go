package slides

import (
	"errors"
	"gh_static_portfolio/internal/ports"
	"net/url"
	"os"
)

type Service struct {
	marpBaseURL string
	slides      ports.SlideRenderer
	pathing     ports.PathingService
	files       ports.FileRepository
}

func New(marpBaseURL string, slides ports.SlideRenderer, pathing ports.PathingService, files ports.FileRepository) *Service {
	return &Service{
		marpBaseURL: marpBaseURL,
		slides:      slides,
		pathing:     pathing,
		files:       files,
	}
}

func (s *Service) GetSlides(nodes ...ports.Node) ([]byte, error) {
	// stat the markdown file
	nodeDir := s.pathing.NodeDirPath(nodes...)
	mdPath := s.pathing.NodeSlidesMarkdownPath(nodes...)
	info, err := os.Stat(mdPath)
	if err != nil {
		return nil, err
	}
	slidesURL, err := url.JoinPath(s.marpBaseURL, mdPath)
	if err != nil {
		return nil, err
	}
	markdownModTime := info.ModTime()

	// stat the html file
	htmlPath := s.pathing.NodeSlidesHTMLPath(nodes...)
	info, err = os.Stat(htmlPath)
	if err != nil {
		// if the html file doesn't exist,
		if errors.Is(err, os.ErrNotExist) {
			// get the html
			content, err := s.slides.NewGetSlides(slidesURL)
			if err != nil {
				return nil, err
			}

			// write to disk, and
			err = s.files.Save(content, nodeDir, "slides.html")
			if err != nil {
				return nil, err
			}

			// return content as string
			return content, nil

		}
		return nil, err
	}
	htmlModTime := info.ModTime()

	// if the html file mod time is older than the markdown file mod time,
	if htmlModTime.Before(markdownModTime) {

		// get the html
		content, err := s.slides.NewGetSlides(slidesURL)
		if err != nil {
			return nil, err
		}

		// write to disk, and
		err = s.files.Update(content, nodeDir, "slides.html")
		if err != nil {
			return nil, err
		}

		// return content as string
		return content, nil
	}

	// if the html file mod time is newer than the markdown file mod time,

	// just return content as string
	content, err := s.files.Read(nodeDir, "slides.html")
	if err != nil {
		return nil, err
	}
	return content, nil

}

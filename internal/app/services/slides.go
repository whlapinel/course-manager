package services

import (
	"errors"
	"fmt"
	"gh_static_portfolio/internal/ports"
	"io/fs"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type SlidesService struct {
	marpBaseURL string
	slides      ports.SlideRenderer
	pathing     ports.PathingService
	files       ports.FileRepository
}

func NewSlidesService(marpBaseURL string, slides ports.SlideRenderer, pathing ports.PathingService, files ports.FileRepository) *SlidesService {
	return &SlidesService{
		marpBaseURL: marpBaseURL,
		slides:      slides,
		pathing:     pathing,
		files:       files,
	}
}

func (s *SlidesService) GetSlides(nodes ...ports.Node) ([]byte, error) {
	log.Println("slidesService.GetSlides running")
	// stat the markdown file
	nodeDir := s.pathing.NodeDirPath(nodes...)
	mdPath := s.pathing.NodeSlidesMarkdownPath(nodes...)
	info, err := os.Stat(mdPath)
	if err != nil {
		// if markdown doesn't exist,
		if errors.Is(err, fs.ErrNotExist) {
			// create the file and return empty content
			err = s.files.Save([]byte{}, nodeDir, filepath.Base(mdPath))
			if err != nil {
				return nil, err
			}
			return []byte{}, nil
		} else {
			return nil, err
		}
	}
	log.Println("slidesPath:", mdPath)
	slidesURL, err := s.dataPathToMarpURL(mdPath)
	if err != nil {
		return nil, err
	}
	log.Println("slidesURL:", slidesURL)

	markdownModTime := info.ModTime()

	// stat the html file
	htmlPath := s.pathing.NodeSlidesHTMLPath(nodes...)
	info, err = os.Stat(htmlPath)
	if err != nil {
		log.Printf("error stat %s: %v", htmlPath, err)
		// if the html file doesn't exist,
		if errors.Is(err, os.ErrNotExist) {
			log.Printf("%s does not exist. fetching from marp server at %s", htmlPath, slidesURL)
			// get the html
			content, err := s.slides.GetSlides(slidesURL)
			if err != nil {
				log.Printf("error fetching from marp server: %s", err)
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
		content, err := s.slides.GetSlides(slidesURL)
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

// for slides editor
func (svc *SlidesService) SlidesContent(nodes ports.Nodes) ([]byte, error) {
	if svc == nil {
		return nil, fmt.Errorf("service is nil")
	}
	if svc.pathing == nil {
		return nil, fmt.Errorf("path service is nil")
	}

	root := svc.pathing.NodeDirPath(nodes.ToSlice()...)
	slidesPath := svc.pathing.NodeSlidesMarkdownPath(nodes.ToSlice()...)
	content, err := svc.files.Read(root, filepath.Base(slidesPath))
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (svc *SlidesService) UpdateSlides(nodes ports.Nodes, content []byte) error {
	root := svc.pathing.NodeDirPath(nodes.ToSlice()...)
	slidesPath := svc.pathing.NodeSlidesMarkdownPath(nodes.ToSlice()...)
	return svc.files.Update(content, root, filepath.Base(slidesPath))
}

func (s *SlidesService) dataPathToMarpURL(path string) (string, error) {
	segments := strings.Split(path, "/")
	return url.JoinPath(s.marpBaseURL, segments[2:]...)
}

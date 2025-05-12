package slides

import (
	"errors"
	"fmt"
	"gh_static_portfolio/internal/ports"
	"net/http"
	"net/url"
	"os"
)

type service struct {
	marpBaseURL string
	paths       ports.PathingService
}

func New(
	marpBaseURL string,
	paths ports.PathingService,
) ports.SlideRenderer {
	return &service{
		marpBaseURL: marpBaseURL,
		paths:       paths,
	}
}

func (s *service) NewGetSlides(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	content := make([]byte, 2048)
	count, err := res.Body.Read(content)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("content of slides was 0")
	}
	return content, nil

}

// GetSlides implements ports.SlideService.
func (s *service) GetSlides(nodes ...ports.Node) ([]byte, error) {
	// stat the markdown file
	mdPath := s.paths.NodeSlidesMarkdownPath(nodes...)
	info, err := os.Stat(mdPath)
	if err != nil {
		return nil, err
	}
	markdownModTime := info.ModTime()

	// stat the html file
	htmlPath := s.paths.NodeSlidesHTMLPath(nodes...)
	info, err = os.Stat(htmlPath)
	if err != nil {
		// if the html file doesn't exist,
		if errors.Is(err, os.ErrNotExist) {
			// get the html
			content, err := s.getSlideContent(nodes...)
			if err != nil {
				return nil, err
			}

			// write to disk, and
			err = s.writeHTMLToFile(content, htmlPath)
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
		content, err := s.getSlideContent(nodes...)
		if err != nil {
			return nil, err
		}

		// write to disk, and
		err = s.writeHTMLToFile(content, htmlPath)
		if err != nil {
			return nil, err
		}

		// return content as string
		return content, nil
	}

	// if the html file mod time is newer than the markdown file mod time,

	// just return content as string
	content, err := s.getSlideContent(nodes...)
	if err != nil {
		return nil, err
	}
	return content, nil

}

func (s *service) getSlideContent(nodes ...ports.Node) ([]byte, error) {
	url, err := s.marpSlidesPath(nodes...)
	if err != nil {
		return nil, err
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	content := make([]byte, 2048)
	count, err := res.Body.Read(content)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("content of slides was 0")
	}
	return content, nil

}

func (s *service) writeHTMLToFile(content []byte, htmlPath string) error {
	err := os.MkdirAll(htmlPath, os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(htmlPath)
	if err != nil {
		return err
	}
	count, err := file.Write(content)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("0 bytes written")
	}
	return nil
}

func (s *service) marpSlidesPath(nodes ...ports.Node) (string, error) {
	relPath := s.paths.NodeDirPath(nodes...)
	fullPath, err := url.JoinPath(s.marpBaseURL, relPath, "slides.md")
	if err != nil {
		return "", err
	}
	return fullPath, nil
}

// // This check to see if the file already exists. if not, should create a new markdown file, write the template to it,
// // and generate the html file. Returns the file path
// func (svc CourseService) CreateSlidesIfNotExist(nodes ...domain.CourseNode) (string, error) {
// 	markdownPath := data.SlidesMarkdownFilePath(nodes...)
// 	_, err := os.Stat(markdownPath)
// 	if os.IsNotExist(err) {
// 		err = os.MkdirAll(filepath.Dir(markdownPath), os.ModePerm)
// 		if err != nil {
// 			return "", err
// 		}
// 		file, err := os.Create(markdownPath)
// 		if err != nil {
// 			return "", err
// 		}
// 		// write template to file
// 		templateFileContents, err := svc.SlidesTemplate(nodes[len(nodes)-1])
// 		if err != nil {
// 			return "", err
// 		}
// 		_, err = file.Write(templateFileContents)
// 		if err != nil {
// 			return "", err
// 		}
// 	}
// 	sitegenerator.GenerateSlides(nodes...)
// 	return markdownPath, nil
// }

// func (svc CourseService) SlidesTemplate(lesson domain.CourseNode) ([]byte, error) {
// 	templateFileContents, err := os.ReadFile("./cmd/web_app/slide_template.md")
// 	if err != nil {
// 		return nil, err
// 	}
// 	templateOther := fmt.Sprintf(
// 		`
// # %s

// # **Warmup**

// # **Agenda**

// # **Looking ahead**`,

// 		lesson.GetName())
// 	templateFileContents = append(templateFileContents, []byte(templateOther)...)
// 	return templateFileContents, nil
// }

// func (svc CourseService) GetSlides(params domain.NodePath) (string, error) {
// 	path, err := marpSlidesPath(params)
// 	if err != nil {
// 		return "", err
// 	}
// 	resp, err := http.Get(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	content := svc.RemoveWatcherScript(body)
// 	return content, nil
// }

// func (svc CourseService) RemoveWatcherScript(html []byte) string {
// 	re := regexp.MustCompile(`(?s)<script>\s*window\.__marpCliWatchWS\s*=\s*"ws://localhost:\d+/[a-f0-9]+";.*?</script>`)
// 	return re.ReplaceAllString(string(html), "")
// }

package marpclient

import (
	"fmt"
	"gh_static_portfolio/internal/ports"
	"io"
	"log"
	"net/http"
)

type marpServerClient struct {
	marpBaseURL string
}

func New(
	marpBaseURL string,
) ports.SlideRenderer {
	return &marpServerClient{
		marpBaseURL: marpBaseURL,
	}
}

func (s *marpServerClient) GetSlides(url string) ([]byte, error) {
	log.Println("client: url:", url)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		content, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			return nil, fmt.Errorf("marp response not ok (%d) and failed to read response body: %w", res.StatusCode, readErr)
		}
		return nil, fmt.Errorf("marp response not ok (%d): %s", res.StatusCode, string(content))
	}
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return content, nil

}

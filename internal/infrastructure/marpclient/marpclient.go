package marpclient

import (
	"bytes"
	"fmt"
	"gh_static_portfolio/internal/ports"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
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
	stripped, err := removeLastScriptFromBody(string(content))
	if err != nil {
		return nil, err
	}
	content = []byte(stripped)
	return content, nil

}

func removeLastScriptFromBody(htmlContent string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	// Find the body element
	var body *html.Node
	var findBody func(*html.Node)
	findBody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			body = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findBody(c)
		}
	}
	findBody(doc)

	if body == nil {
		return htmlContent, nil // No body found, return original
	}

	// Find the last script element within body
	var lastScript *html.Node
	var findLastScript func(*html.Node)
	findLastScript = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "script" {
			lastScript = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findLastScript(c)
		}
	}
	findLastScript(body)

	// Remove the last script if found
	if lastScript != nil {
		lastScript.Parent.RemoveChild(lastScript)
	}

	// Render back to HTML
	var buf bytes.Buffer
	err = html.Render(&buf, doc)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

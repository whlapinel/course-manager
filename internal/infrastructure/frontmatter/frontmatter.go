package frontmatter

import (
	"bytes"
	"fmt"
	"gh_static_portfolio/internal/ports"

	"github.com/BurntSushi/toml"
)

type frontMatter struct {
}

func New() ports.FrontMatterReadWriter {
	return &frontMatter{}
}

func (f *frontMatter) ParseFrontMatter(data []byte) (ports.MarkdownFile, error) {
	var file ports.MarkdownFile

	// Check for TOML front matter
	// if !bytes.HasPrefix(data, []byte("+++")) {
	// 	return nil, string(data), fmt.Errorf("no TOML front matter found")
	// }

	if !bytes.HasPrefix(data, []byte("+++")) {
		file.Content = string(data)
		return file, nil
	}

	// Split on +++
	parts := bytes.SplitN(data, []byte("+++"), 3)
	if len(parts) < 3 {
		return file, fmt.Errorf("invalid front matter format")
	}

	var fm ports.FrontMatter
	err := toml.Unmarshal(bytes.TrimSpace(parts[1]), &fm)
	if err != nil {
		return file, err
	}

	content := string(bytes.TrimLeft(parts[2], "\n\r"))
	file.FrontMatter = &fm
	file.Content = content

	return file, nil
}

func (f *frontMatter) ToBytes(file ports.MarkdownFile) ([]byte, error) {
	var buf bytes.Buffer

	// Write opening delimiter
	buf.WriteString("+++\n")

	// Encode front matter
	encoder := toml.NewEncoder(&buf)
	if err := encoder.Encode(file.FrontMatter); err != nil {
		return nil, err
	}

	// Write closing delimiter and content
	buf.WriteString("+++\n\n")
	buf.WriteString(file.Content)
	return buf.Bytes(), nil
}

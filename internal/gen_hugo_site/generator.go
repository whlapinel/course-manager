package genhugosite

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Generate(repo data.CourseRepo) error {
	err := os.RemoveAll("./internal/my_site/content")
	if err != nil {
		return err
	}
	terms, err := repo.GetTerms()
	if err != nil {
		return err
	}
	for _, term := range terms {
		path := NodeDirPath(term)
		workingDir := "./internal/my_site"
		err = os.Chdir(workingDir)
		if err != nil {
			return err
		}
		cmd := exec.Command("hugo", "new", "content", path)
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil

}

func NodeDirPath(nodes ...domain.CourseNode) string {
	var path = "content"
	for _, node := range nodes {
		path = filepath.Join(path, strings.ToLower(node.TypeName()+"s"))
		path = filepath.Join(
			path,
			fmt.Sprintf("%s-%d", strings.ToLower(node.TypeName()), node.GetID()),
		)
		path = filepath.Join(path, "_index.md")
	}
	return path
}

package genhugosite

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

func Generate(repo data.CourseRepo) error {
	contentPath := "./internal/my_site/content"
	err := os.RemoveAll(contentPath)
	if err != nil {
		return err
	}
	publicPath := "./internal/my_site/public"
	err = os.RemoveAll(publicPath)
	if err != nil {
		return err
	}
	err = os.MkdirAll(contentPath, os.ModePerm)
	if err != nil {
		return err
	}
	terms, err := repo.GetTerms()
	if err != nil {
		return err
	}
	var nodes []domain.CourseNode
	for _, term := range terms {
		nodes = append(nodes, term)
	}
	err = GenerateHugo(nodes)
	if err != nil {
		return err
	}
	for _, term := range terms {
		courses, err := repo.GetCourses(term.ID)
		if err != nil {
			return err
		}
		var nodes []domain.CourseNode
		for _, course := range courses {
			nodes = append(nodes, course)
		}
		err = GenerateHugo(nodes, term)
		if err != nil {
			return err
		}
		for _, course := range courses {
			units, err := repo.GetUnits(course.ID)
			if err != nil {
				return err
			}
			var nodes []domain.CourseNode
			for _, unit := range units {
				nodes = append(nodes, unit)
			}
			err = GenerateHugo(nodes, term, course)
			if err != nil {
				return err
			}
			for _, unit := range units {
				lessons, err := repo.GetLessons(unit.ID)
				if err != nil {
					return err
				}
				var nodes []domain.CourseNode
				for _, lesson := range lessons {
					nodes = append(nodes, lesson)
				}
				err = GenerateHugo(nodes, term, course, unit)
				if err != nil {
					return err
				}

			}
		}
	}
	return nil

}

func GenerateHugo(children []domain.CourseNode, parents ...domain.CourseNode) error {
	listParents := append(parents, children[0])
	listPath := NodeListDirPath(listParents...)
	listIndexPath := BranchBundlePage(listPath)
	err := WriteNodeListPageToMarkdown(children[0], listIndexPath)
	// err := CreateHugoContent(listIndexPath)
	if err != nil {
		return err
	}
	for _, child := range children {
		nodePath := NodeDirPath(listPath, child)
		nodePath = BranchBundlePage(nodePath)
		err := WriteNodePageToMarkdown(child, nodePath)
		// err = CreateHugoContent(nodePath)
		if err != nil {
			return err
		}
	}
	return nil

}

func BranchBundlePage(path string) string {
	return filepath.Join(path, "_index.md")
}

func NodeListDirPath(nodes ...domain.CourseNode) string {
	var path = "./internal/my_site/content"
	for i, node := range nodes {
		path = filepath.Join(path, strings.ToLower(node.TypeName()+"s"))
		if i == len(nodes)-1 {
			break
		}
		path = filepath.Join(
			path,
			fmt.Sprintf("%s-%d", strings.ToLower(node.TypeName()), node.GetID()),
		)
	}
	return path
}

func NodeDirPath(listPath string, node domain.CourseNode) string {
	path := filepath.Join(
		listPath,
		fmt.Sprintf("%s-%d", strings.ToLower(node.TypeName()), node.GetID()),
	)
	return path
}

func WriteNodeListPageToMarkdown(node domain.CourseNode, path string) error {
	tpl := template.Must(template.ParseFiles("internal/gen_hugo_site/node_list.md"))
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error in creating file: %v", err)
	}
	data := NodeListPageData{
		Date: time.Now().Format("2006-01-02T15:04:05-07:00"),
		Node: node,
	}
	err = tpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("error in executing template: %v", err)
	}
	return nil

}

func WriteNodePageToMarkdown(node domain.CourseNode, path string) error {
	tpl := template.Must(template.ParseFiles("internal/gen_hugo_site/node.md"))
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error in creating file: %v", err)
	}
	data := NodePageData{
		Date:                 time.Now().Format("2006-01-02T15:04:05-07:00"),
		Node:                 node,
		SanitizedDescription: strings.ReplaceAll(node.GetDescription(), "\n", " "),
	}
	err = tpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("error in executing template: %v", err)
	}
	return nil
}

type NodePageData struct {
	Date                 string
	Node                 domain.CourseNode
	SanitizedDescription string
}

type NodeListPageData struct {
	Date string
	Node domain.CourseNode
}

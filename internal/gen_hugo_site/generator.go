package genhugosite

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"time"
)

func Generate(repo data.CourseRepo) error {
	contentPath := "./internal/my_site/content"
	publicPath := "./internal/my_site/public"

	// Cleanup old content
	if err := os.RemoveAll(contentPath); err != nil {
		return err
	}
	if err := os.RemoveAll(publicPath); err != nil {
		return err
	}
	if err := os.MkdirAll(contentPath, os.ModePerm); err != nil {
		return err
	}

	terms, err := repo.GetTerms()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 100) // Buffered channel for error handling

	// Process terms in parallel
	for _, term := range terms {
		wg.Add(1)
		go func(term domain.CourseNode) {
			defer wg.Done()
			if err := processTerm(repo, term, errChan); err != nil {
				errChan <- err
			}
		}(term)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)

	// Collect any errors
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func processTerm(repo data.CourseRepo, term domain.CourseNode, errChan chan error) error {
	err := GenerateHugo([]domain.CourseNode{term})
	if err != nil {
		return err
	}

	courses, err := repo.GetCourses(term.GetID())
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, course := range courses {
		wg.Add(1)
		go func(course domain.CourseNode) {
			defer wg.Done()
			if err := processCourse(repo, term, course, errChan); err != nil {
				errChan <- err
			}
		}(course)
	}
	wg.Wait()
	return nil
}

func processCourse(repo data.CourseRepo, term, course domain.CourseNode, errChan chan error) error {
	err := GenerateHugo([]domain.CourseNode{course}, term)
	if err != nil {
		return err
	}

	units, err := repo.GetUnits(course.GetID())
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, unit := range units {
		wg.Add(1)
		go func(unit domain.CourseNode) {
			defer wg.Done()
			if err := processUnit(repo, term, course, unit, errChan); err != nil {
				errChan <- err
			}
		}(unit)
	}
	wg.Wait()
	return nil
}

func processUnit(repo data.CourseRepo, term, course, unit domain.CourseNode, errChan chan error) error {
	err := GenerateHugo([]domain.CourseNode{unit}, term, course)
	if err != nil {
		return err
	}

	lessons, err := repo.GetLessons(unit.GetID())
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, lesson := range lessons {
		wg.Add(1)
		go func(lesson domain.CourseNode) {
			defer wg.Done()
			if err := GenerateHugo([]domain.CourseNode{lesson}, term, course, unit); err != nil {
				errChan <- err
			}
		}(lesson)
	}
	wg.Wait()
	return nil
}

func GenerateHugo(children []domain.CourseNode, parents ...domain.CourseNode) error {
	listParents := append(parents, children[0])
	listPath := NodeListDirPath(listParents...)
	listIndexPath := BranchBundlePage(listPath)
	var parentNode domain.CourseNode
	if parents != nil {
		parentNode = parents[len(parents)-1]
	} else {
		parentNode = &domain.User{
			FirstName: "Billy",
			LastName:  "Bob",
		}
	}

	// Use Mutex to prevent concurrent file writing issues
	var mu sync.Mutex
	mu.Lock()
	err := WriteNodeListPageToMarkdown(children[0], parentNode, listIndexPath)
	mu.Unlock()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(children))

	for _, child := range children {
		wg.Add(1)
		go func(child domain.CourseNode) {
			defer wg.Done()
			nodePath := NodeDirPath(listPath, child)
			nodePath = BranchBundlePage(nodePath)

			mu.Lock()
			err := WriteNodePageToMarkdown(child, nodePath)
			mu.Unlock()

			if err != nil {
				errChan <- err
			}
		}(child)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
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

func WriteNodeListPageToMarkdown(node, parentNode domain.CourseNode, path string) error {
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
		Date:       time.Now().Format("2006-01-02T15:04:05-07:00"),
		Node:       node,
		ParentName: parentNode.GetName(),
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
	Date       string
	Node       domain.CourseNode
	ParentName string
}

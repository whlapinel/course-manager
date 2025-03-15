package service

import (
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	templates "gh_static_portfolio/internal/templates/static"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const AssetsDir = "./sites/assets"

func NodeListPageURL(nodes ...domain.CourseNode) string {
	path := NodeListPagePath(nodes...)
	return URL(path)
}

func NodeDetailsPageURL(nodes ...domain.CourseNode) string {
	path := NodeDetailsPagePath(nodes...)
	return URL(path)
}

// intended as a callback for breadcrumbs
func NodeURL(nodes ...domain.CourseNode) string {
	currNode := nodes[len(nodes)-1]
	if currNode.ChildTypeName() == "" {
		return NodeDetailsPageURL(nodes...)
	} else {
		return NodeListPageURL(nodes...)
	}
}

// remove sites/{username} from path for URL
func URL(path string) string {
	segments := strings.SplitN(path, "/", 3)
	if len(segments) > 2 {
		return "/" + segments[2] // Keep everything after the first two segments
	}
	return "/" // Return root if there aren't enough segments
}

func FilePath(relPath string, nodes ...domain.CourseNode) string {
	path := NodeFilesPath(nodes...)
	path = filepath.Join(path, relPath)
	return path
}

func FileURL(relPath string, nodes ...domain.CourseNode) string {
	return URL(FilePath(relPath, nodes...))
}

func ViewMarkdownPath(relPath string, nodes ...domain.CourseNode) string {
	path := FilePath(relPath, nodes...)
	path = strings.ReplaceAll(path, ".md", ".html")
	return path
}

func ViewMarkdownURL(relPath string, nodes ...domain.CourseNode) string {
	return URL(ViewMarkdownPath(relPath, nodes...))
}

func CourseCalendarURL(nodes ...domain.CourseNode) string {
	return URL(CourseCalendarPagePath(nodes...))
}

func StaticAssetsURL(relPath string) string {
	path := filepath.Join("/sites/assets", relPath)
	return URL(path)
}

func StaticSiteRootDir(user domain.User) string {
	return filepath.Join("sites", user.Username())
}

func StaticNodePath(nodes ...domain.CourseNode) string {
	user, ok := nodes[0].(domain.User)
	if !ok {
		panic("node is not user")
	}
	path := StaticSiteRootDir(user)
	for i, node := range nodes[1:] {
		if i != 0 {
			path = strings.ToLower(filepath.Join(path, node.TypeName()+"s"))
		}
		path = strings.ToLower(filepath.Join(
			path,
			fmt.Sprintf("%s_%d", node.TypeName(), node.GetID()),
		))

	}
	return path
}

func CourseCalendarPagePath(nodes ...domain.CourseNode) string {
	path := StaticNodePath(nodes...)
	leafNode := nodes[len(nodes)-1] // last node is current node
	path = filepath.Join(path, fmt.Sprintf("%s_%d_calendar.html", strings.ToLower(leafNode.TypeName()), leafNode.GetID()))
	return path
}

func NodeListPagePath(nodes ...domain.CourseNode) string {
	path := StaticNodePath(nodes...)
	leafNode := nodes[len(nodes)-1] // last node is current node
	path = filepath.Join(path, fmt.Sprintf("%s_%d.html", strings.ToLower(leafNode.TypeName()), leafNode.GetID()))
	return path
}

func NodeDetailsPagePath(nodes ...domain.CourseNode) string {
	path := StaticNodePath(nodes...)
	currNode := nodes[len(nodes)-1] // last node is current node
	path = filepath.Join(path, fmt.Sprintf("%s_%d_details.html", strings.ToLower(currNode.TypeName()), currNode.GetID()))
	return path
}

// Student-facing site
func NodeImagePath(nodes ...domain.CourseNode) string {
	dir := StaticNodePath(nodes...)
	return filepath.Join(dir, "image.png")
}

func NodeFilesPath(nodes ...domain.CourseNode) string {
	dir := StaticNodePath(nodes...)
	return filepath.Join(dir, "files")
}

func NodeSlidesPath(nodes ...domain.CourseNode) string {
	dir := StaticNodePath(nodes...)
	return filepath.Join(dir, "slides.html")
}

func NodeDetailsPage(nodes domain.Nodes) templates.StaticNodeDetailsPage {
	return templates.NewStaticNodeDetailsPage(templates.StaticNodeDetailsParams{
		Node:     nodes.CurrentNode(),
		Path:     NodeDetailsPagePath(nodes.ToSlice()...),
		PageData: Layout(nodes),
	})
}

func NodeListPage(nodes domain.Nodes) (templates.StaticNodeListPage, error) {
	page, err := templates.NewStaticNodeListPage(templates.StaticNodeListParams{
		PageData:                 Layout(nodes),
		Nodes:                    nodes,
		ListChildChildrenURLFunc: NodeListPageURL,
		ChildDetailsURLFunc:      NodeDetailsPageURL,
		Path:                     NodeListPagePath(nodes.ToSlice()...),
		CourseCalendarURL:        CourseCalendarURL,
	})
	if err != nil {
		return templates.StaticNodeListPage{}, err
	}
	return page, nil
}

func Layout(nodes domain.Nodes) templates.PageData {
	return templates.PageData{
		User:        nodes.User,
		AssetsURL:   StaticAssetsURL,
		BreadCrumbs: templates.BreadCrumbs(nodes, NodeURL),
	}
}

func (svc CourseService) generate(user domain.User, term domain.Term) error {
	homePage := templates.NewHomePage(templates.StaticHomePageParams{
		HomePage: templates.HomePage{
			Path: filepath.Join(StaticSiteRootDir(user), "index.html"),
			PageData: templates.PageData{
				User:      user,
				AssetsURL: StaticAssetsURL,
			},
		},
	})
	errChan := make(chan error, 1000)
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	var ctx, cancel = context.WithCancel(context.Background())

	wg.Add(1)
	go RenderPage(homePage, errChan, wg, cancel)
	nodes := domain.Nodes{
		User: user,
		Term: term,
	}
	err := renderNodePages(nodes, errChan, wg, ctx, cancel)
	if err != nil {
		return err
	}
	for _, course := range term.Courses {
		nodes := domain.Nodes{
			User:   user,
			Term:   term,
			Course: course,
		}
		cc := templates.NewCourseCalendarPage(templates.NewCourseCalendarParams{
			CourseCalendarPage: templates.CourseCalendarPage{
				Nodes:         nodes,
				Path:          CourseCalendarPagePath(nodes.ToSlice()...),
				AssetsURL:     StaticAssetsURL,
				LessonPageURL: NodeDetailsPageURL,
			},
		})
		svc.renderMarkdownFiles(course.GetName(), data.NodeFilesDirPath(nodes.ToSlice()...))
		svc.renderMarkdownFiles(course.GetName(), NodeFilesPath(nodes.ToSlice()...))
		wg.Add(1)
		go RenderPage(cc, errChan, wg, cancel)
		err := data.CopyNodeDir(data.NodeDirPath(nodes.ToSlice()...), StaticNodePath(nodes.ToSlice()...))
		if err != nil {
			return err
		}
		err = renderNodePages(nodes, errChan, wg, ctx, cancel)
		if err != nil {
			return err
		}
		for _, unit := range course.Units {
			nodes := domain.Nodes{
				User:   user,
				Term:   term,
				Course: course,
				Unit:   unit,
			}
			svc.renderMarkdownFiles(unit.Designation(), data.NodeFilesDirPath(nodes.ToSlice()...))
			svc.renderMarkdownFiles(unit.Designation(), NodeFilesPath(nodes.ToSlice()...))
			err = renderNodePages(nodes, errChan, wg, ctx, cancel)
			if err != nil {
				return err
			}
			for _, lesson := range unit.Lessons {
				nodes := domain.Nodes{
					User:   user,
					Term:   term,
					Course: course,
					Unit:   unit,
					Lesson: lesson,
				}
				svc.renderMarkdownFiles(lesson.Designation(), data.NodeFilesDirPath(nodes.ToSlice()...))
				svc.renderMarkdownFiles(lesson.Designation(), NodeFilesPath(nodes.ToSlice()...))
				renderNodePages(nodes, errChan, wg, ctx, cancel)
			}
		}
	}
	go func() {
		wg.Wait()
		close(errChan)
	}()
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func renderNodePages(nodes domain.Nodes, errChan chan<- error, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) error {
	detailsPage := NodeDetailsPage(nodes)
	var staticPage templates.StaticPage
	if nodes.CurrentNode().TypeName() == domain.LessonTypeName.String() {
		lessonpage := templates.StaticLessonDetailsPage{
			Nodes:                 nodes,
			StaticNodeDetailsPage: detailsPage,
			LessonSlidesURL:       NodeSlidesPath(nodes.ToSlice()...),
			ViewMarkdownURL:       ViewMarkdownURL,
			FilesURLFunc:          FileURL,
		}
		staticPage = lessonpage
	} else {
		staticPage = detailsPage
	}
	select {
	case <-ctx.Done():
		return nil
	default:
		wg.Add(1)
		go RenderPage(staticPage, errChan, wg, cancel)
		if nodes.CurrentNode().ChildTypeName() != "" {
			listPage, err := NodeListPage(nodes)
			if err != nil {
				return err
			}
			wg.Add(1)
			go RenderPage(listPage, errChan, wg, cancel)

		}
	}
	return nil
}

func RenderPage(page templates.StaticPage, errChan chan<- error, wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()
	err := os.MkdirAll(filepath.Dir(page.Filepath()), os.ModePerm)
	if err != nil {
		errChan <- err
		cancel()
		return
	}
	file, err := os.Create(page.Filepath())
	if err != nil {
		errChan <- err
		cancel()
		return
	}
	defer file.Close()
	err = page.Component().Render(context.Background(), file)
	if err != nil {
		errChan <- err
		cancel()
		return
	}
}

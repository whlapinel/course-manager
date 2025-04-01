package service

import (
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	templates "gh_static_portfolio/internal/templates/static"
	"io"
	"io/fs"
	"log"
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

func NodeFilesPagePath(relPath string, nodes ...domain.CourseNode) string {
	path := FilePath(relPath, nodes...)
	path = filepath.Join(path, "files.html")
	return path
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
	path := NodeFilesRootPath(nodes...)
	path = filepath.Join(path, relPath)
	return path
}

func FilesPageURL(relPath string, nodes ...domain.CourseNode) string {
	return URL(NodeFilesPagePath(relPath, nodes...))

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
	if len(nodes) < 2 {
		return path
	}
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

func NodeFilesRootPath(nodes ...domain.CourseNode) string {
	dir := StaticNodePath(nodes...)
	return filepath.Join(dir, "files")
}

// file copied from data folder, simply marp written to a file
func NodeSlidesPath(nodes ...domain.CourseNode) string {
	dir := StaticNodePath(nodes...)
	return filepath.Join(dir, "slides.html")
}

// iframe with src set to node slides path
func SlidesFragmentPath(nodes ...domain.CourseNode) string {
	dir := StaticNodePath(nodes...)
	return filepath.Join(dir, "slides-fragment.html")
}
func AssessmentsFragmentPath(nodes ...domain.CourseNode) string {
	dir := StaticNodePath(nodes...)
	return filepath.Join(dir, "assessment-fragment.html")
}

func SlidesFragment(nodes domain.Nodes) templates.StaticPage {
	return templates.SlidesFragment{
		Path:            SlidesFragmentPath(nodes.ToSlice()...),
		LessonSlidesURL: URL(NodeSlidesPath(nodes.ToSlice()...)),
	}
}

func AssessmentsFragment(page templates.StaticLessonDetailsPage, assessments []domain.Assessment, nodes domain.Nodes) templates.StaticPage {
	return templates.AssessmentsFragment{
		StaticLessonDetailsPage: page,
		Path:                    AssessmentsFragmentPath(nodes.ToSlice()...),
		Assessments:             assessments,
	}
}

func NodeDetailsPage(nodes domain.Nodes) templates.StaticNodeDetailsPage {
	return templates.StaticNodeDetailsPage{
		Node:         nodes.CurrentNode(),
		Path:         NodeDetailsPagePath(nodes.ToSlice()...),
		PageData:     Layout(nodes),
		FilesPageURL: FilesPageURL("", nodes.ToSlice()...),
	}
}

func NodeFilesPages(relPath string, nodes domain.Nodes) []templates.FilesPageSection {
	var pages []templates.FilesPageSection
	path := FilePath(relPath, nodes.ToSlice()...)
	files, _ := os.ReadDir(path)
	var filesList []templates.File
	for _, file := range files {
		if file.Name() == filepath.Base(NodeFilesPagePath(relPath, nodes.ToSlice()...)) {
			continue
		}
		subRelPath := filepath.Join(relPath, file.Name())
		item := templates.File{
			Name:  file.Name(),
			Path:  filepath.Join(path, file.Name()),
			URL:   FileURL(subRelPath, nodes.ToSlice()...),
			IsDir: file.IsDir(),
		}
		if item.IsDir {
			item.URL = FilesPageURL(subRelPath, nodes.ToSlice()...)
			pages = append(pages, NodeFilesPages(subRelPath, nodes)...)
		}
		filesList = append(filesList, item)
	}
	var dirName string
	if relPath == "" {
		dirName = "Root"
	} else {
		dirName = relPath
	}
	parentDir := filepath.Dir(relPath)
	if parentDir == "." {
		parentDir = "Root"
	}
	page := templates.FilesPageSection{
		Root: relPath == "",
		Path: NodeFilesPagePath(relPath, nodes.ToSlice()...),
		ParentDirectory: templates.File{
			Name:  parentDir,
			URL:   URL(NodeFilesPagePath(filepath.Dir(relPath), nodes.ToSlice()...)),
			Path:  FilePath(filepath.Dir(relPath), nodes.ToSlice()...),
			IsDir: true,
		},
		CurrentDirectory: templates.File{
			Path: FilePath(relPath, nodes.ToSlice()...),
			Name: dirName,
			URL:  FileURL(relPath, nodes.ToSlice()...),
		},
		Files: filesList,
	}
	pages = append(pages, page)
	return pages

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
	nodes := domain.Nodes{
		User: user,
		Term: term,
	}
	err := os.RemoveAll(StaticNodePath(nodes.User))
	if err != nil {
		return err
	}
	err = CopyNodeDir(data.NodeDirPath(nodes.ToSlice()...), StaticNodePath(nodes.ToSlice()...))
	if err != nil {
		return err
	}
	homePage := templates.NewHomePage(templates.StaticHomePageParams{
		HomePage: templates.HomePage{
			Path: filepath.Join(StaticSiteRootDir(user), "index.html"),
			PageData: templates.PageData{
				User:      user,
				AssetsURL: StaticAssetsURL,
			},
			Term:    term,
			TermURL: NodeListPageURL(nodes.ToSlice()...),
		},
	})
	errChan := make(chan error, 1000)
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	var ctx, cancel = context.WithCancel(context.Background())

	wg.Add(1)
	go RenderPage(homePage, errChan, wg, cancel)
	err = renderNodePages(nodes, errChan, wg, ctx, cancel)
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
		wg.Add(1)
		go RenderPage(cc, errChan, wg, cancel)
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
				renderLessonPages(nodes, errChan, wg, ctx, cancel)
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

// any file or directory with this prefix will not be copied to the static site files
const IgnorePrefix = "secret"

// for copying all files in the lesson, unit, or course directory
// should not copy anything that begins with "secret"
func CopyNodeDir(srcRoot, destRoot string) error {
	log.Println("copying", srcRoot, "to", destRoot)
	// if directory doesn't exist, early return
	_, err := os.Stat(srcRoot)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}
	files, err := os.ReadDir(srcRoot)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return nil
	}
	return filepath.WalkDir(srcRoot, func(srcPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(srcRoot, srcPath)
		if err != nil {
			return err
		}
		name := filepath.Base(relPath)
		if strings.HasPrefix(name, IgnorePrefix) || strings.HasSuffix(name, IgnorePrefix) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		destPath := filepath.Join(destRoot, relPath)
		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}
		return copyFile(srcPath, destPath)
	})
}

func copyFile(srcPath, destPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	return err
}

// renders details page and list page for each node
func renderNodePages(nodes domain.Nodes, errChan chan<- error, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) error {
	staticPage := NodeDetailsPage(nodes)
	select {
	case <-ctx.Done():
		return nil
	default:
		wg.Add(1)
		go RenderPage(staticPage, errChan, wg, cancel)
		filesPages := NodeFilesPages("", nodes)
		for _, page := range filesPages {
			wg.Add(1)
			go RenderPage(page, errChan, wg, cancel)
		}
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

// renders details page and list page for each node
func renderLessonPages(nodes domain.Nodes, errChan chan<- error, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) error {
	nodePage := NodeDetailsPage(nodes)
	if nodePage.AssetsURL == nil {
		log.Fatal("AssetsURL func is nil!!")
	}
	lessonpage := templates.StaticLessonDetailsPage{
		Nodes:                 nodes,
		StaticNodeDetailsPage: nodePage,
		LessonSlidesURL:       URL(SlidesFragmentPath(nodes.ToSlice()...)),
		AssessmentsURL:        URL(AssessmentsFragmentPath(nodes.ToSlice()...)),
		ViewMarkdownURL:       ViewMarkdownURL,
		FilesURLFunc:          FileURL,
	}
	select {
	case <-ctx.Done():
		return nil
	default:
		wg.Add(1)
		go RenderPage(lessonpage, errChan, wg, cancel)
		slidesFrag := SlidesFragment(nodes)
		wg.Add(1)
		go RenderPage(slidesFrag, errChan, wg, cancel)
		assFrag := AssessmentsFragment(lessonpage, nodes.Lesson.Assessments, nodes)
		wg.Add(1)
		go RenderPage(assFrag, errChan, wg, cancel)
		filesPages := NodeFilesPages("", nodes)
		for _, page := range filesPages {
			wg.Add(1)
			go RenderPage(page, errChan, wg, cancel)
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
	log.Println("creating file: ", page.Filepath())
	file, err := os.Create(page.Filepath())
	if err != nil {
		log.Println("error creating file:", err)
		errChan <- err
		cancel()
		return
	}
	defer file.Close()
	err = page.Component().Render(context.Background(), file)
	if err != nil {
		log.Println("error rendering html to file: ", page.Filepath())
		errChan <- err
		cancel()
		return
	}
}

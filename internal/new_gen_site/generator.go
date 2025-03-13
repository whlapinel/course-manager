package newgensite

import (
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/new_gen_site/templates"
	"os"
	"path/filepath"
	"strings"
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
	// should skip term and start with courses, which is third node
	if len(nodes) < 3 {
		return "/index.html"
	}
	for _, node := range nodes[2:] {
		path = strings.ToLower(filepath.Join(path, node.TypeName()+"s"))
		path = strings.ToLower(filepath.Join(
			path,
			fmt.Sprintf("%s_%d", node.TypeName(), node.GetID()),
		))

	}
	return path
}

func CoursesListPagePath(user domain.User) string {
	path := StaticSiteRootDir(user)
	return filepath.Join(path, "courses.html")
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

func NodeDetailsPage(user domain.User, nodes domain.Nodes) templates.StaticNodeDetailsPage {
	return templates.NewStaticNodeDetailsPage(templates.StaticNodeDetailsParams{
		Node:         nodes.CurrentNode(),
		Path:         NodeDetailsPagePath(nodes.ToSlice()...),
		StaticLayout: Layout(user, nodes),
	})
}

func NodeListPage(user domain.User, nodes domain.Nodes) (templates.StaticNodeListPage, error) {
	page, err := templates.NewStaticNodeListPage(templates.StaticNodeListParams{
		StaticLayout:             Layout(user, nodes),
		Nodes:                    nodes,
		ListChildChildrenURLFunc: NodeListPageURL,
		ChildDetailsURLFunc:      NodeDetailsPageURL,
		Path:                     NodeListPagePath(nodes.ToSlice()...),
	})
	if err != nil {
		return templates.StaticNodeListPage{}, err
	}
	return page, nil
}

func Layout(user domain.User, nodes domain.Nodes) templates.StaticLayout {
	return templates.StaticLayout{
		User:        user,
		AssetsURL:   StaticAssetsURL,
		BreadCrumbs: templates.BreadCrumbs(nodes, NodeURL),
	}
}

func Generate(user domain.User, term domain.Term) error {
	nodes := domain.Nodes{
		User: user,
		Term: term,
	}
	courseListPage, err := NodeListPage(user, nodes)
	courseListPage.Path = CoursesListPagePath(user) // need to override default path
	if err != nil {
		return err
	}
	err = RenderPage(courseListPage)
	if err != nil {
		return err
	}
	homePage := templates.NewHomePage(templates.StaticHomePageParams{
		HomePage: templates.HomePage{
			Path: filepath.Join(StaticSiteRootDir(user), "index.html"),
			StaticLayout: templates.StaticLayout{
				User:      user,
				AssetsURL: StaticAssetsURL,
			},
		},
	})
	err = RenderPage(homePage)
	if err != nil {
		return err
	}
	for _, course := range term.Courses {
		nodes := domain.Nodes{
			User:   user,
			Term:   term,
			Course: course,
		}
		err = data.CopyNodeDir(data.NodeDirPath(nodes.ToSlice()...), StaticNodePath(nodes.ToSlice()...))
		if err != nil {
			return err
		}
		courseDetailsPage := NodeDetailsPage(user, nodes)
		err = RenderPage(courseDetailsPage)
		if err != nil {
			return err
		}
		unitListPage, err := NodeListPage(user, nodes)
		if err != nil {
			return err
		}
		err = RenderPage(unitListPage)
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
			unitDetailsPage := NodeDetailsPage(user, nodes)
			err = RenderPage(unitDetailsPage)
			if err != nil {
				return err
			}
			lessonListPage, err := NodeListPage(user, nodes)
			if err != nil {
				return err
			}
			err = RenderPage(lessonListPage)
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
				nodes.Lesson = lesson
				lessonDetailsPage := NodeDetailsPage(user, nodes)
				err = RenderPage(lessonDetailsPage)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func RenderPage(page templates.StaticPage) error {
	err := os.MkdirAll(filepath.Dir(page.Filepath()), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(page.Filepath())
	if err != nil {
		return err
	}
	err = page.Component().Render(context.Background(), file)
	if err != nil {
		return err
	}
	return nil

}

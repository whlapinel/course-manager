package newgensite

import (
	"context"
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"gh_static_portfolio/internal/new_gen_site/templates"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const AssetsDir = "./sites/assets"

func NodePageURL(nodes ...domain.CourseNode) string {
	path := NodePagePath(nodes...)
	log.Println("NodePageURL: path:", path)
	return URL(path)
}

// remove sites/{username} from path for URL
func URL(path string) string {
	segments := strings.SplitN(path, "/", 3)
	if len(segments) > 2 {
		return "/" + segments[2] // Keep everything after the first two segments
	}
	return "/" // Return root if there aren't enough segments
}

func StaticAssetsPath(relPath string) string {
	path := filepath.Join("/sites", relPath)
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

func NodePagePath(nodes ...domain.CourseNode) string {
	leafNode := nodes[len(nodes)-1]
	path := StaticNodePath(nodes...)
	path = filepath.Join(path, fmt.Sprintf("%s_%d.html", strings.ToLower(leafNode.TypeName()), leafNode.GetID()))
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

func Generate(user domain.User, term domain.Term) error {
	nodes := domain.Nodes{
		User: user,
		Term: term,
	}
	courseListPage, err := templates.NewStaticNodeListPage(templates.StaticNodeListParams{
		Nodes:        nodes,
		ChildUrlFunc: NodePageURL,
		Path:         CoursesListPagePath(user),
	})
	if err != nil {
		return err
	}
	err = RenderPage(courseListPage)
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
		unitListPage, err := templates.NewStaticNodeListPage(templates.StaticNodeListParams{
			Nodes:        nodes,
			ChildUrlFunc: NodePageURL,
			Path:         NodePagePath(nodes.ToSlice()...),
		})
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
			log.Println("nodes.Lesson:", nodes.Lesson.GetName())
			lessonListPage, err := templates.NewStaticNodeListPage(templates.StaticNodeListParams{
				Nodes:        nodes,
				ChildUrlFunc: NodePageURL,
				Path:         NodePagePath(nodes.ToSlice()...),
			})
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
				lessonDetailsPage := templates.NewStaticNodeDetailsPage(templates.StaticNodeDetailsParams{
					Node: nodes.CurrentNode(),
					Path: NodePagePath(nodes.ToSlice()...),
				})
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

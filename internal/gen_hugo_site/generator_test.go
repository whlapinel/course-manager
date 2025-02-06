package genhugosite

import (
	"fmt"
	"gh_static_portfolio/internal/data"
	"gh_static_portfolio/internal/domain"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var repo data.CourseRepo

func TestMain(m *testing.M) {
	rootDir, _ := filepath.Abs("../../")
	fmt.Println("Setting test working directory to:", rootDir)

	err := os.Chdir(rootDir)
	if err != nil {
		fmt.Println("Chdir failed:", err)
		os.Exit(1)
	}
	queries, db, err := data.InitDB("internal/data/database/course_manager.db")
	if err != nil {
		os.Exit(1)
	}
	defer db.Close()
	repo = data.NewCourseRepo(queries)

	os.Exit(m.Run())
}

func TestGenerate(t *testing.T) {
	err := Generate(repo)
	if err != nil {
		t.Error(err)
	}

}

func TestWriteNodeToMarkdown(t *testing.T) {
	node, err := repo.GetTermByID(1)
	if err != nil {
		t.Error(err)
	}
	err = WriteNodeListPageToMarkdown(node, &domain.User{}, "./internal/gen_hugo_site/test_list_node.md")
	if err != nil {
		t.Error(err)
	}
	err = WriteNodePageToMarkdown(node, "./internal/gen_hugo_site/test_node.md")
	if err != nil {
		t.Error(err)
	}

}

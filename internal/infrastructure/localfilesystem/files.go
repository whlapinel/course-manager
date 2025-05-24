package localfilesystem

import (
	"fmt"
	"gh_static_portfolio/internal/ports"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type localFileSystem struct {
}

func New() ports.FileRepository {
	return &localFileSystem{}
}

// returns FileInfo. Create rootDir if it doesn't exist
func (l *localFileSystem) FileInfo(relPath string, rootDir string) (ports.FileInfo, error) {
	fileInfo := ports.FileInfo{}
	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		return ports.FileInfo{}, err
	}
	root, err := os.OpenRoot(rootDir)
	if err != nil {
		return ports.FileInfo{}, err
	}
	relPath = filepath.Clean(relPath)
	defer root.Close()
	info, err := root.Stat(relPath)
	if err != nil {
		return fileInfo, err
	}
	if info.IsDir() {
		fileInfo.IsDir = true
	}
	absRoot, err := filepath.Abs(rootDir)
	if err != nil {
		return fileInfo, err
	}
	absPath := filepath.Join(absRoot, relPath)
	fileInfo.AbsPath = absPath
	return fileInfo, nil
}

// same as Save but without check to make sure the file doesn't already exist
func (l *localFileSystem) Update(contents []byte, rootDir, relPath string) error {
	// Open the root directory as an os.Root
	root, err := os.OpenRoot(rootDir)
	if err != nil {
		return fmt.Errorf("failed to open root directory: %w", err)
	}
	defer root.Close()
	// Ensure the parent directories exist
	dir := filepath.Dir(relPath)
	log.Println("dir", dir)
	if dir != "." {
		if err := root.Mkdir(dir, os.ModePerm); err != nil && !os.IsExist(err) {
			return fmt.Errorf("failed to create directories: %w", err)
		}
	}
	// Create the newFile within the root
	newFile, err := root.Create(relPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer newFile.Close()
	// Write the contents to the file
	if _, err := newFile.Write(contents); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

// Delete implements file.FileRepository.
// if directory is not empty it will return *PathError
func (l *localFileSystem) Delete(rootDir, relPath string) error {
	root, err := os.OpenRoot(rootDir)
	if err != nil {
		return fmt.Errorf("failed to open root directory: %w", err)
	}
	defer root.Close()
	err = root.Remove(relPath)
	if err != nil {
		return err
	}
	return nil

}

// List implements file.FileRepository.
func (l *localFileSystem) List(rootDir, relPath string) ([]fs.DirEntry, error) {
	root, err := os.OpenRoot(rootDir)
	if err != nil {
		return nil, fmt.Errorf("failed to open root directory: %w", err)
	}
	defer root.Close()
	info, err := root.Stat(relPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("file is not a directory: %w", err)
	}
	system := root.FS()
	files, err := fs.ReadDir(system, relPath)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// Read implements file.FileRepository.
func (l *localFileSystem) Read(rootDir, relPath string) ([]byte, error) {
	root, err := os.OpenRoot(rootDir)
	if err != nil {
		return nil, fmt.Errorf("failed to open root directory: %w", err)
	}
	defer root.Close()
	openedFile, err := root.Open(relPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", relPath, err)
	}
	defer openedFile.Close()
	data, err := io.ReadAll(openedFile)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (l *localFileSystem) Save(contents []byte, rootDir, relPath string) error {
	// Open the root directory as an os.Root
	err := os.MkdirAll(rootDir, os.ModePerm)
	if err != nil {
		return err
	}
	root, err := os.OpenRoot(rootDir)
	if err != nil {
		return fmt.Errorf("failed to open root directory: %w", err)
	}
	defer root.Close()
	// Ensure the parent directories exist
	dir := filepath.Dir(relPath)
	if dir != "." {
		if err := root.Mkdir(dir, 0755); err != nil && !os.IsExist(err) {
			return fmt.Errorf("failed to create directories: %w", err)
		}
	}
	_, err = root.Stat(relPath)
	if err == nil {
		return ErrFileExists
	}
	// Create the newFile within the root
	newFile, err := root.Create(relPath)
	if err != nil {
		return fmt.Errorf("failed to create file in %s: %w", relPath, err)
	}
	defer newFile.Close()
	// Write the contents to the file
	if _, err := newFile.Write(contents); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

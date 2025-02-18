package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// copyFile copies a file from src to dst, creating necessary directories
func copyFile(src, dst string) error {
	// Create the destination directory if it doesn't exist
	err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		return err
	}

	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the file contents
	_, err = io.Copy(dstFile, srcFile)
	return err
}

// searchAndCopy searches recursively for files with the prefix "secret_" and copies them
func searchAndCopy(srcDir, dstDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Check if the file has the prefix "secret_"
		if strings.HasPrefix(info.Name(), "secret_") {
			// Determine the new destination path, preserving the directory structure
			relPath, err := filepath.Rel(srcDir, path)
			if err != nil {
				return err
			}

			dstPath := filepath.Join(dstDir, relPath)

			// Copy the file
			fmt.Printf("Copying %s to %s\n", path, dstPath)
			if err := copyFile(path, dstPath); err != nil {
				return err
			}
		}

		return nil
	})
}

func main() {
	srcDir := "internal/data/users/user_101602110272674353046" // Replace with actual source directory
	dstDir := "secret_files"                                   // Replace with actual destination directory

	err := searchAndCopy(srcDir, dstDir)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Copy operation completed successfully.")
	}
}

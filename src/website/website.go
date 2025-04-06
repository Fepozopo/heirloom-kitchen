package website

import (
	"io"
	"os"
	"path/filepath"
)

// CopyStaticToPublic copies all the contents of the static directory to the public directory recursively.
func CopyStaticToPublic() error {
	staticPath := "static"
	publicPath := "public"

	// Remove all contents of the public directory if it exists
	if _, err := os.Stat(publicPath); err == nil {
		err = os.RemoveAll(publicPath)
		if err != nil {
			return err
		}
	}

	// Create the public directory if it doesn't exist
	err := os.MkdirAll(publicPath, os.ModePerm)
	if err != nil {
		return err
	}

	// Walk through the static directory and copy files to the public directory
	err = filepath.Walk(staticPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Determine the target path in the public directory
		relPath, err := filepath.Rel(staticPath, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(publicPath, relPath)

		if info.IsDir() {
			// Create directories in the public directory
			return os.MkdirAll(targetPath, os.ModePerm)
		} else {
			// Copy files to the public directory
			return copyFile(path, targetPath)
		}
	})

	return err
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	return err
}

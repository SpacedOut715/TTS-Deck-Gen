package imageprocessing

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Thank you ChatGPT!!!

// findEndDirectories traverses the root directory and returns a slice of paths to end directories.
func findEndDirectories(root string) ([]string, error) {
	var endDirs []string

	// Walk the directory tree.
	err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip if it's not a directory
		if !info.IsDir() {
			return nil
		}

		// Check if it's an end directory
		isEndDir, err := isEndDirectory(path)
		if err != nil {
			return err
		}

		if isEndDir {
			endDirs = append(endDirs, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return endDirs, nil
}

// isEndDirectory checks if the given directory is an end directory (no subdirectories).
func isEndDirectory(dir string) (bool, error) {
	// Open the directory.
	d, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer d.Close()

	// Read directory contents.
	contents, err := d.Readdir(-1) // Read all directory contents.
	if err != nil {
		return false, err
	}

	// Check if there are any subdirectories.
	for _, entry := range contents {
		if entry.IsDir() && entry.Name() != "." && entry.Name() != ".." {
			return false, nil
		}
	}

	return true, nil
}

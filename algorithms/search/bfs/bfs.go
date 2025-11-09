// Breadth-First Search (BFS) - File Finder
// Time: O(n) - where n is total number of files and directories
// Space: O(w) - where w is maximum width (number of directories at one level)

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFileBFS(directoryName, fileName string) string {
	// Initialize the queue with starting directory
	queue := []string{directoryName}

	for len(queue) > 0 {
		// Take the first folder from the queue (FIFO)
		dir := queue[0]
		queue = queue[1:]
		// Read the folder contents
		files, _ := os.ReadDir(dir)

		// Go through every item inside the current folder
		for _, item := range files {
			fullPath := filepath.Join(dir, item.Name())

			// Check if it's a file or a folder
			if item.IsDir() {
				// If it's a folder — add it to the queue to explore later
				queue = append(queue, fullPath)
			} else if strings.Contains(item.Name(), fileName) {
				// If it's a file and matches our search — return it
				return fullPath
			}
		}
	}
	return ""
}

func main() {
	result := FindFileBFS("directoryName", "fileName")
	if result != "" {
		fmt.Println(result)
	}
}

/*
How BFS file search works:

Explores directories level by level (breadth-first traversal).
Uses a queue to defer subdirectory exploration until all items at current level are checked.

Parameters:
  - directoryName: root folder to start search
  - fileName: name or part of the file name to search for

Returns:
  - full path to the found file, or empty string if not found

Used Go standard library:

os.ReadDir(dir)
  - Reads the contents of a directory
  - Returns []fs.DirEntry and an error

filepath.Join(dir, file)
  - Joins path segments into a normalized path string
  - Example: filepath.Join("pics", "a.png") → "pics/a.png"

item.Name()
  - Method of fs.DirEntry interface
  - Returns the name of the file or directory

item.IsDir()
  - Method of fs.DirEntry interface
  - Returns true if the entry is a directory, false if it's a file

strings.Contains(str, substr)
  - Checks whether substr is within str
  - Returns true if substr is found, false otherwise
*/

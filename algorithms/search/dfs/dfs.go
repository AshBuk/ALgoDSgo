// Depth-First Search (DFS) - File Finder
// Time: O(n) - where n is total number of files and directories
// Space: O(h) - where h (height) is maximum depth of directory tree (call stack for recursion)

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindFileDFS(directoryName, fileName string) string {
	files, _ := os.ReadDir(directoryName)

	for _, file := range files {
		fullPath := filepath.Join(directoryName, file.Name())
		if file.IsDir() {
			if result := FindFileDFS(fullPath, fileName); result != "" {
				return result
			}
		} else if strings.Contains(file.Name(), fileName) {
			return fullPath
		}

	}
	return ""
}

func main() {
	result := FindFileDFS("directoryName", "fileName")
	if result != "" {
		fmt.Println(result)
	}
}

/*
How DFS file search works:

Explores each directory fully before moving to siblings (depth-first traversal).
Recursively descends into subdirectories immediately when encountered.

Parameters:
  - directoryName: path to the directory to traverse
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

file.Name()
  - Method of fs.DirEntry interface
  - Returns the name of the file or directory

file.IsDir()
  - Method of fs.DirEntry interface
  - Returns true if the entry is a directory, false if it's a file

strings.Contains(str, substr)
  - Checks whether substr is within str
  - Returns true if substr is found, false otherwise
*/

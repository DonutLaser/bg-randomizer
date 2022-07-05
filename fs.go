package main

import (
	"fmt"
	"io/fs"
	"os"
)

func ReadDir(path string) ([]fs.DirEntry, bool) {
	dir, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []fs.DirEntry{}, false
	}
	defer dir.Close()

	items, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return []fs.DirEntry{}, false
	}

	return items, true
}

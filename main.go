package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/reujab/wallpaper"
)

func main() {
	basePath := "D:/Wallpapers"
	files, success := ReadDir(basePath)
	if !success {
		return
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(files))

	fullPath := fmt.Sprintf("%s/%s", basePath, files[index].Name())
	err := wallpaper.SetFromFile(fullPath)
	if err != nil {
		fmt.Println(err)
	}
}

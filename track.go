package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	// "database/sql"
	//_ "github.com/mattn/go-sqlite3"
)

func getAllTracks() []Track {
	var result []Track
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	dir = dir + "/Music/Vocaloid/emonloid4"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}

	for i, file := range files {
		if isAudio(file) {
			track := Track{
				Name:   file.Name(),
				Artist: "emonloid4",
                Album: "Cool",
				Number: i,
			}
			result = append(result, track)
		}
	}

	return result
}

type Track struct {
	Name   string
	Artist string
    Album  string
	Number int
}

// AI
func isAudio(file fs.DirEntry) bool {
	fileName := file.Name()
	extension := strings.ToLower(filepath.Ext(fileName))

	switch extension {
	case ".mp3", ".wav", ".flac":
		return true
	default:
		return false
	}
}

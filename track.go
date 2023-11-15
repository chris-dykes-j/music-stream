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

func getAllTracks() []string {
    var result []string
    dir, err := os.UserHomeDir()
    if err != nil {
        fmt.Println(err)
    }
    dir = dir + "/Music/Vocaloid/emonloid4"
    files, err := os.ReadDir(dir)
    if err != nil {
        fmt.Println(err)
    }

    for _, file := range files {
        if isAudio(file) {
            result = append(result, file.Name())
        }
    }
    
    return result 
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

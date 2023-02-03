// Package watch monitors markdown files for changes and generates html

package watch

import (
	"log"
	"os"
	"strings"

	"github.com/edwinwalela/tiwi/pkg/parse"
	"github.com/fsnotify/fsnotify"
)

// getMdFileName extracts markdown file name from directory path
func getMdFileName(path string) string {
	vals := strings.Split(path, "/")
	for i := range vals {
		if strings.Contains(vals[i], ".md") {
			return vals[i]
		}
	}
	return ""
}

func Watch(args []string) {
	var projectDir string
	wDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get working directory: %s", err.Error())
	}
	if len(args) == 0 {
		projectDir = wDir
	} else {
		projectDir = wDir + "/" + args[0]
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("failed to watch directory: %s", err.Error())
	}
	defer watcher.Close()
	parse.Build(args, false, true)

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Printf("Event occurred: %v\n", getMdFileName(event.Name))

			case err := <-watcher.Errors:
				log.Fatalf("failed to watch directory: %s", err.Error())
			}
		}
	}()
	if err := watcher.Add(projectDir); err != nil {
		log.Fatalf("Failed to watch project directory: %s", err.Error())
	}

	<-done
}

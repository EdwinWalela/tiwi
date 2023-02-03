// Package watch monitors markdown files for changes and generates html

package watch

import (
	"log"
	"os"
	"strings"

	"github.com/edwinwalela/tiwi/pkg/parse"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

// blue defines colored formatted output
var blue = color.New(color.FgCyan).PrintfFunc()

// green defines colored formatted output
var green = color.New(color.FgGreen).PrintfFunc()

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

func logEvent(event *fsnotify.Event) {
	log.Printf("File changed: %v. Rebuilding HTML\n", getMdFileName(event.Name))
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
	parse.Build(args, false, true, false)

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				logEvent(&event)
				parse.Build(args, false, true, true)
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

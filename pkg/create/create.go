// Package create generates the tiwi project by creating the project folder
// and default markdown files

package create

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/enescakir/emoji"
	"github.com/fatih/color"
)

// defaultPages defines a mapping of default markdown files to their content
var defaultPages = map[string]string{
	"index": `# Tiwi
![tiwi-logo](https://raw.githubusercontent.com/EdwinWalela/tiwi/main/docs/tiwi-mini.png)

---

Generate HTML from markdown

[Getting Started](./getting-started.md)

[Sample Article](./first-article.md)

---

[View on Github](https://github.com/EdwinWalela/tiwi)
			`,
	"first-article": `[Home Page](./index.md)
# This is my first article
This is a sample article
	`,
	"getting-started": `[Home Page](./index.md)
# Getting Started
---
## Supported markdown syntax
- #- h1
- ##- h2 
- ###- h3
- ![alt](src)- img
- --- - horizontal separator 
- []() - link

## Generate markdown
Run ./tiwi build my-site
`,
	"not-found": `
# Page not found

The page requested was not found.

[Home Page](./index.md)
	`,
}

// createParentFolder creates a folder for the project
func createParentFolder(title string) error {
	if _, err := os.Stat(title); os.IsNotExist(err) {
		if err := os.Mkdir(title, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create project directory: %v", err.Error())
		}
	}
	return nil
}

// createPages creates the default markdown files for a new project
func createPages(parentDir string) error {
	for title := range defaultPages {
		if err := createPage(title, parentDir); err != nil {
			return fmt.Errorf("failed to create page %s: %v", title, err.Error())
		}
	}
	return nil
}

// createPage writes markdown to file for a new project
func createPage(page string, parentDir string) error {
	md := []byte(defaultPages[page])

	if err := os.WriteFile(fmt.Sprintf("%s/%s.md", parentDir, page), md, 0644); err != nil {
		return err
	}
	return nil
}

// CreateSite generates a new tiwi project with its respective default markdown files
func CreateSite(args []string) {
	startTime := time.Now()
	projectTitle := ""
	blue := color.New(color.FgCyan).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()

	if len(args) > 0 {
		projectTitle = strings.Join(args, "-")
	} else if len(args) != 0 {
		projectTitle = args[0]
	} else {
		fmt.Printf("\nPlease specify the project name:\n\n")
		blue("tiwi create ")
		green("<my-project>\n\n")
		fmt.Printf("For example:\n\n")
		blue("tiwi create ")
		green("my-tiwi-site\n\n")
		os.Exit(1)
	}
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err.Error())
	}

	fmt.Printf("\n%v Scaffolding project in :", emoji.Star)
	green("%s/%s...\n", currentDir, projectTitle)
	if err := createParentFolder(projectTitle); err != nil {
		log.Fatal(err.Error())
	}
	if err := createPages(projectTitle); err != nil {
		log.Fatal(err.Error())
	}
	timeDiff := time.Since(startTime)
	fmt.Printf("\n%v Project created in %s", emoji.ThumbsUp, timeDiff.Round(time.Millisecond))
	if projectTitle != "." {

		fmt.Printf("\n\ncd ")
		blue("%s\n", projectTitle)
	} else {
		fmt.Println("")
	}
	fmt.Printf("\nTo build project run: ")
	blue("tiwi build\n\n")

	fmt.Printf("\n%v Start building!\n\n", emoji.PersonSurfing)
}

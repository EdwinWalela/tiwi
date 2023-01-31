package create

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
)

var DefaultPages = map[string]string{
	"index": `
# Tiwi

Generate HTML from markdown

## Getting Started

### Supported markdown syntax

- #- h1
- ##- h2 
- ###- h3
- ![alt](src)- img
- --- - div
- []() - link


### Generate HTML

./tiwi build


[View on Github](https://github.com/EdwinWalela/tiwi)
	`,
	"not-found": `
# Page not found

The page requested was not found.
	`,
}

func createParentFolder(title string) error {
	if _, err := os.Stat(title); os.IsNotExist(err) {
		if err := os.Mkdir(title, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create project directory: %v", err.Error())
		}
	}
	return nil
}

func createPages(parentDir string) error {
	for title := range DefaultPages {
		if err := createPage(title, parentDir); err != nil {
			return fmt.Errorf("failed to create page %s: %v", title, err.Error())
		}
	}
	return nil
}

func createPage(page string, parentDir string) error {
	about := []byte(fmt.Sprintf("%s", DefaultPages[page]))

	if err := os.WriteFile(fmt.Sprintf("%s/%s.md", parentDir, page), about, 0644); err != nil {
		return err
	}
	return nil
}

func CreateSite(args []string) {
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

	fmt.Printf("\nScaffolding project in :")
	green("%s/%s...\n", currentDir, projectTitle)
	if err := createParentFolder(projectTitle); err != nil {
		log.Fatal(err.Error())
	}
	if err := createPages(projectTitle); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("\nProject created")
	if projectTitle != "." {

		fmt.Printf("\n\ncd ")
		blue("%s\n", projectTitle)
	} else {
		fmt.Println("")
	}
	fmt.Printf("\nTo build project run: ")
	blue("tiwi build\n\n")

	fmt.Printf("\nStart building!\n\n")

}

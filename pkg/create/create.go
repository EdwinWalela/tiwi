package create

import (
	"fmt"
	"log"
	"os"
)

var DefaultPages = []string{"index", "not-found"}

func createParentFolder(title string) error {
	if _, err := os.Stat(title); os.IsNotExist(err) {
		if err := os.Mkdir(title, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create project directory: %v", err.Error())
		}
	}
	return nil
}

func createPages(parentDir string) error {
	for _, page := range DefaultPages {
		if err := createPage(page, parentDir); err != nil {
			return fmt.Errorf("failed to create page %s: %v", page, err.Error())
		}
	}
	return nil
}

func createPage(page string, parentDir string) error {
	about := []byte(fmt.Sprintf("# %s\n", page))

	if err := os.WriteFile(fmt.Sprintf("%s/%s.md", parentDir, page), about, 0644); err != nil {
		return err
	}
	return nil
}

func CreateSite(args []string) {
	projectTitle := ""
	if len(args) != 0 {
		projectTitle = args[0]
	}
	currentDir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err.Error())
	}

	fmt.Printf("\nScaffolding project in : %s/%s...\n", currentDir, projectTitle)
	if err := createParentFolder(projectTitle); err != nil {
		log.Fatal(err.Error())
	}
	if err := createPages(projectTitle); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("\nProject created")
	if projectTitle != "." {

		fmt.Printf("\n\ncd %s\n", projectTitle)
	} else {
		fmt.Println("")
	}
	fmt.Printf("\nStart building!\n\n")
}

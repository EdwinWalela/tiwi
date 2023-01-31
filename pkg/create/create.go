package create

import (
	"fmt"
	"log"
	"os"
)

var DefaultPages = []string{"index", "not-found"}

func createParentFolder() {
	if _, err := os.Stat("site"); os.IsNotExist(err) {
		if err := os.Mkdir("site", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

func createPages() {
	for _, page := range DefaultPages {
		if err := createPage(page); err != nil {
			log.Fatalf("Failed to create page %s: %v", page, err.Error())
		}
	}
}

func createPage(page string) error {
	about := []byte(fmt.Sprintf("# %s\n", page))

	if err := os.WriteFile(fmt.Sprintf("site/%s.md", page), about, 0644); err != nil {
		return err
	}
	return nil
}

func CreateSite() {
	createParentFolder()
	createPages()
}

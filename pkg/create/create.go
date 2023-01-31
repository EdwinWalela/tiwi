package create

import (
	"log"
	"os"
)

func createParentFolder() {
	if _, err := os.Stat("site"); os.IsNotExist(err) {
		if err := os.Mkdir("site", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

func createPages() {
	createPage()
}

func createPage() {
	about := []byte("# About\n")

	if err := os.WriteFile("site/about.md", about, 0644); err != nil {
		log.Fatal(err)
	}
}

func CreateSite() {
	createParentFolder()
}

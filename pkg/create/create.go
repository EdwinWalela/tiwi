package create

import (
	"log"
	"os"
)

func CreateSite() {
	if err := os.Mkdir("site", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	about := []byte("# About\n")

	if err := os.WriteFile("/site/about.md", about, 0644); err != nil {
		log.Fatal(err)
	}

}

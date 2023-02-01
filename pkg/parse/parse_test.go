package parse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPages(t *testing.T) {
	pages, err := getPages("../../test-site")
	assert.NoError(t, err)
	assert.Equal(t, 2, len(pages))
}

func TestCreateOutputFolder(t *testing.T) {
	directory := "../../test-site"
	err := createOutputFolder(directory)

	assert.NoError(t, err)
	assert.DirExists(t, directory+"/static")
}

func TestReadMarkDown(t *testing.T) {
	md, err := readMarkdown("index.md", "../../test-site")
	assert.NoError(t, err)
	expectedMd := `
# Tiwi
![tiwi-logo](https://raw.githubusercontent.com/EdwinWalela/tiwi/main/docs/tiwi-mini.png)

---

Generate HTML from markdown

## Getting Started

### Supported markdown syntax

- #- h1
- ##- h2 
- ###- h3
- ![alt](src)- img
- --- - div
- []() - link

---

### Generate HTML

./tiwi build

---

[View on Github](https://github.com/EdwinWalela/tiwi)
	`
	assert.Equal(t, expectedMd, md)
}

func TestGenerateAnchorTag(t *testing.T) {
	expectedHtml := "<a href=\"link\">title</a>"
	targetMd := "[title](link)"
	html := parseMd(targetMd)

	assert.Equal(t, expectedHtml, html)

}

func TestWriteHTML(t *testing.T) {

}

func TestBuildProject(t *testing.T) {
	args := []string{"../../test-site"}
	Build(args)
	assert.DirExists(t, fmt.Sprintf("%s/static", args[0]))
}

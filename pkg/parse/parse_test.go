package parse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPages(t *testing.T) {
	pages, err := getPages("../../test-site")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(pages))
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
	expectedMd := `# Tiwi
![tiwi-logo](https://raw.githubusercontent.com/EdwinWalela/tiwi/main/docs/tiwi-mini.png)

---

Generate HTML from markdown

[Getting Started](./getting-started.md)

[Sample Article](./first-article.md)

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

func TestGenerateImgTag(t *testing.T) {
	expectedHtml := "<img alt=\"alt\" src=\"link\"/>"
	targetMd := "![alt](link)"
	html := parseMd(targetMd)

	assert.Equal(t, expectedHtml, html)
}

func TestGenerateHTMLElement(t *testing.T) {
	expectedHtml := []string{
		"<h1>heading 1</h1>",
		"<h2>heading 2</h2>",
		"<h3>heading 3</h3>",
		"<div style=\"border-top:solid 1px gray\"></div>",
	}

	markdown := []string{
		"# heading 1",
		"## heading 2",
		"### heading 3",
		"---",
	}

	for i, md := range markdown {
		html := parseMd(md)
		assert.Equal(t, expectedHtml[i], html)
	}
}

func TestWriteHTML(t *testing.T) {

}

func TestBuildProject(t *testing.T) {
	args := []string{"../../test-site"}
	Build(args, false)
	assert.DirExists(t, fmt.Sprintf("%s/static", args[0]))
}

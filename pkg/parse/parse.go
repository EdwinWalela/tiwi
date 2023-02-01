// Package parse reads markdown file and generates HTML from the markdown

package parse

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"strings"

	"github.com/enescakir/emoji"
	"github.com/fatih/color"
)

// mdTohtml defines mappings between markdown and HTML opening tags
var mdTohtml = map[string]string{
	"#":   "<h1>",
	"##":  "<h2>",
	"###": "<h3>",
	"[":   "<a href=\"%s\">%s</a>",
	"!":   "<img alt=\"%s\" src=\"%s\"/>",
	"---": "<div style=\"border-top:solid 1px gray\"></div>",
}

// htmlOpenToClose defines mappings between HTML opening and closing tags
var htmlOpenToClose = map[string]string{
	"<h1>": "</h1>",
	"<h2>": "</h2>",
	"<h3>": "</h3>",
}

// htmlHeader defines the default HTML header to be used for the generated HTML files
var htmlHeader string = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="UTF-8" />
	<head>
		<title>%s</title>
	</head>
`

// htmlBody defines the default HTML body element to be used for the generated HTML files
var htmlBody string = `
	<body>
	%s
	</body>
 </html>
`

// blue defines colored formatted output
var blue = color.New(color.FgCyan).PrintfFunc()

// green defines colored formatted output
var green = color.New(color.FgGreen).PrintfFunc()

// parseAnchorTag generates HTML anchor tags from markdown
func parseAnchorTag(src string) string {
	title, link, _ := strings.Cut(src, "]")
	link = strings.ReplaceAll(link, "(", "")
	link = strings.ReplaceAll(link, ")", "")
	title = strings.ReplaceAll(title, "[", "")

	return fmt.Sprintf(mdTohtml[src[0:1]], link, title)
}

// parseImgTag generates HTML img elements from markdown
func parseImgTag(src string) string {
	// v -> entire line of markdown
	// src -> markdown syntax e.g # ## ###
	imgAlt, imgSrc, _ := strings.Cut(src, "]")

	imgSrc = strings.ReplaceAll(imgSrc, "(", "")
	imgSrc = strings.ReplaceAll(imgSrc, ")", "")

	imgAlt = strings.ReplaceAll(imgAlt, "[", "")
	imgAlt = strings.ReplaceAll(imgAlt, "!", "")
	return fmt.Sprintf(mdTohtml[src[0:1]], imgAlt, imgSrc)
}

// createOutputFolder generates a folder called static where generated HTML files are saved
func createOutputFolder(projectDir string) error {
	path := "./static"
	if projectDir != "" {
		path = fmt.Sprintf("%s/static", projectDir)
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err.Error())
		}
	}
	return nil
}

// writeHTML writes the generated HTML to file
func writeHTML(src string, target string, projectDir string) {
	staticPath := "static"
	if projectDir != "" {
		staticPath = fmt.Sprintf("%v/static", projectDir)
	}
	data := []byte(fmt.Sprintf(htmlHeader, "title") + fmt.Sprintf(htmlBody, src))
	if err := os.WriteFile(fmt.Sprintf("%s/%s.html", staticPath, strings.ReplaceAll(target, ".md", "")), data, 0644); err != nil {
		log.Fatal(err.Error())
	}
}

// getPages locates and returns a list of markdown file names in the current or specified directory
func getPages(projectDir string) ([]string, error) {
	pages := []string{}
	path := "./"
	if projectDir != "" {
		path = projectDir
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if strings.Contains(f.Name(), ".md") {
			pages = append(pages, f.Name())
		}
	}
	return pages, nil
}

// readMarkdown reads markdown from file and returns it as a string
func readMarkdown(page string, projectDir string) (string, error) {
	pagePath := page
	if projectDir != "" {
		pagePath = fmt.Sprintf("%s/%s", projectDir, page)
	}
	dat, err := os.ReadFile(pagePath)

	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %s", page, err.Error())
	}
	return string(dat), nil
}

// parseMd generates HTML from a markdown string
func parseMd(md string) string {
	html := ""
	el, val, _ := strings.Cut(md, " ")

	if _, exists := mdTohtml[el]; !exists {
		if len(el) > 0 {
			if el[0:1] == "[" {
				html = parseAnchorTag(md)
			} else if el[0:1] == "!" {
				html = parseImgTag(md)
			} else {
				html = "<p>" + md + "</p>"
			}
		}
	} else {
		html = mdTohtml[el] + val + htmlOpenToClose[mdTohtml[el]]
	}
	return html
}

// buildHTML generates html files from markdown
func buildHTML(src string, page string, projectDir string) error {
	html := ""
	for _, v := range strings.Split(src, "\n") {
		if v == "" {
			continue
		}
		html += "\t\t" + parseMd(v) + "\n"
	}
	writeHTML(html, page, projectDir)
	return nil
}

// Build reads markdown files and generates HTML files
func Build(args []string) {
	var projectDir string
	if len(args) != 0 {
		projectDir = args[0]
	}

	pages, err := getPages(projectDir)

	if err != nil {
		log.Fatal(err.Error())
	}

	if len(pages) == 0 {
		log.Fatalf("No Markdown files were found in the directory\n")
	}

	fmt.Printf("\n%v Found", emoji.PageFacingUp)
	green(" [%d] ", len(pages))
	fmt.Printf("page(s):\n\n")
	for _, page := range pages {
		green("- %s\n", page)
	}

	if err := createOutputFolder(projectDir); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%v Generating HTML...\n", emoji.HourglassNotDone)

	for _, page := range pages {
		md, err := readMarkdown(page, projectDir)
		if err != nil {
			log.Fatalf("failed to read markdown file %s: %s", page, err.Error())
		}
		if err := buildHTML(md, page, projectDir); err != nil {
			log.Fatal(err.Error())
		}
	}

	blue("\n%v Process complete.", emoji.ThumbsUp)
	fmt.Printf(" HTML files generated at ")
	outputPath := "./static"
	if projectDir != "" {
		outputPath = fmt.Sprintf("%s/static", projectDir)
	}
	green("%s\n\n", outputPath)
}

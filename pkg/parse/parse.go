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

var mdTohtml = map[string]string{
	"#":   "<h1>",
	"##":  "<h2>",
	"###": "<h3>",
	"[":   "<a href=\"%s\">%s</a>",
	"!":   "<img alt=\"%s\" src=\"%s\"/>",
	"---": "<div style=\"border-top:solid 1px gray\"</div>",
}

var htmlOpenToClose = map[string]string{
	"<h1>": "</h1>",
	"<h2>": "</h2>",
	"<h3>": "</h3>",
}

var htmlHeader string = `
	<head>
		<title>%s</title>
	</head>
`

var htmlBody string = `
 <body>
 %s
 </body>
`

func parseAnchorTag(src string) string {
	title, link, _ := strings.Cut(src, "]")
	fmt.Println("src=" + src)
	fmt.Println("title=" + title)
	fmt.Println("link=" + link)
	link = strings.ReplaceAll(link, "(", "")
	link = strings.ReplaceAll(link, ")", "")
	title = strings.ReplaceAll(title, "[", "")

	return fmt.Sprintf(mdTohtml[src[0:1]], link, title)
}

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

func Build(args []string) {
	var projectDir string
	if len(args) != 0 {
		projectDir = args[0]
	}
	blue := color.New(color.FgCyan).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()
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

	if err := createOutputFolder(projectDir); err != nil {
		log.Fatal(err)
	}

	for _, page := range pages {
		green("- %s\n", page)
	}

	fmt.Printf("\n%v Generating HTML...\n", emoji.HourglassNotDone)

	for _, page := range pages {
		pagePath := page
		if projectDir != "" {
			pagePath = fmt.Sprintf("%s/%s", projectDir, page)
		}
		dat, err := os.ReadFile(pagePath)

		if err != nil {
			log.Fatal(err.Error())
		}
		src := string(dat)
		html := ""

		for _, v := range strings.Split(src, "\n") {
			if v != "" {
				el, val, _ := strings.Cut(v, " ")
				_ = val

				if _, exists := mdTohtml[el]; !exists {

					if len(el) > 0 {
						if el[0:1] == "[" {
							html += "\t\t" + parseAnchorTag(v) + "\n"
						} else if el[0:1] == "!" {
							html += "\t\t" + parseImgTag(v) + "\n"
						} else {
							html += "\t\t<p>" + v + "</p>" + "\n"
						}
					}
				} else {
					html += "\t\t" + mdTohtml[el] + val + htmlOpenToClose[mdTohtml[el]] + "\n"
				}
			}
		}
		writeHTML(html, page, projectDir)
	}

	blue("\n%v Process complete.", emoji.ThumbsUp)
	fmt.Printf(" HTML files generated at ")
	outputPath := "./static"
	if projectDir != "" {
		outputPath = fmt.Sprintf("%s/static", projectDir)
	}
	green("%s\n\n", outputPath)
}

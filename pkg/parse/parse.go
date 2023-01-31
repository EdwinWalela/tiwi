package parse

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"strings"

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

	link = strings.ReplaceAll(link, "(", "")
	link = strings.ReplaceAll(link, ")", "")
	title = strings.ReplaceAll(title, "[", "")

	return fmt.Sprintf(mdTohtml[src[0:1]], link, title)
}

func parseImgTag(src string, v string) string {
	// v -> entire line of markdown
	// src -> markdown syntax e.g # ## ###
	imgAlt, imgSrc, _ := strings.Cut(v, "]")

	imgSrc = strings.ReplaceAll(imgSrc, "(", "")
	imgSrc = strings.ReplaceAll(imgSrc, ")", "")

	imgAlt = strings.ReplaceAll(imgAlt, "[", "")
	imgAlt = strings.ReplaceAll(imgAlt, "!", "")
	return fmt.Sprintf(mdTohtml[src[0:1]], imgAlt, imgSrc)
}

func createOutputFolder() error {
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		if err := os.Mkdir("static", os.ModePerm); err != nil {
			return fmt.Errorf("failed to create output directory: %v", err.Error())
		}
	}
	return nil
}

func writeHTML(src string, target string) {
	data := []byte(fmt.Sprintf(htmlHeader, "title") + fmt.Sprintf(htmlBody, src))
	if err := os.WriteFile(fmt.Sprintf("static/%s.html", strings.ReplaceAll(target, ".md", "")), data, 0644); err != nil {
		log.Fatal(err.Error())
	}
}

func getPages() ([]string, error) {
	pages := []string{}
	files, err := ioutil.ReadDir("./")
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

func Build() {

	blue := color.New(color.FgCyan).PrintfFunc()
	green := color.New(color.FgGreen).PrintfFunc()
	pages, err := getPages()
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(pages) == 0 {
		log.Fatalf("No Markdown files were found in the directory\n")
	}

	fmt.Printf("\nFound")
	green(" [%d] ", len(pages))
	fmt.Printf("page(s):\n\n")

	if err := createOutputFolder(); err != nil {
		log.Fatal(err)
	}

	for _, page := range pages {
		green("- %s\n", page)
	}

	fmt.Println("\nGenerating HTML...")

	for _, page := range pages {
		dat, err := os.ReadFile(page)

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
							html += parseAnchorTag(el) + "\n"
						} else if el[0:1] == "!" {
							html += parseImgTag(el, v) + "\n"
						} else {
							html += "<p>" + v + "</p>" + "\n"
						}
					}
				} else {
					html += mdTohtml[el] + val + htmlOpenToClose[mdTohtml[el]] + "\n"
				}
			}
		}
		writeHTML(html, page)
	}

	blue("\nProcess complete.")
	fmt.Printf(" HTML files generated at ")
	green("./static\n\n")
}

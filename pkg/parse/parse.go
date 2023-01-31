package parse

import (
	"fmt"
	"log"
	"os"
	"strings"
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

func Build() {
	dat, err := os.ReadFile("site/index.md")
	if err != nil {
		log.Fatal(err.Error())
	}
	src := string(dat)

	for _, v := range strings.Split(src, "\n") {
		if v != "" {
			el, val, _ := strings.Cut(v, " ")
			_ = val

			if _, exists := mdTohtml[el]; !exists {
				if el[0:1] == "[" {
					fmt.Println(parseAnchorTag(el))
				} else if el[0:1] == "!" {
					fmt.Println(parseImgTag(el, v))
				} else {
					fmt.Println("<p>", v, "</p>")
				}
			} else {
				fmt.Println(mdTohtml[el], val, htmlOpenToClose[mdTohtml[el]])
			}
		}

	}
}

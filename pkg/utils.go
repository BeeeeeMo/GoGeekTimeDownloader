package pkg

import (
	"encoding/base64"
	"os"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	log "github.com/sirupsen/logrus"
)

func CleanFileName(name string) string {
	regex := regexp.MustCompile(`[\\/:*?"<>|]`)
	return regex.ReplaceAllString(name, "")
}

func WriteToFile(path string, content string) {
	f, err := os.Create(path)
	if err != nil {
		log.Error(err)
		return
	}
	defer f.Close()
	f.WriteString(content)
}

func HtmlToMarkdown(html string) string {
	converter := md.NewConverter("", true, nil)
	markdown, _ := converter.ConvertString(html)
	// regex to find all the image urls
	regex := regexp.MustCompile(`!\[.*\]\((.*)\)`)
	// enumerate all the image urls
	for _, url := range regex.FindAllStringSubmatch(markdown, -1) {
		img := Fetch(url[1])
		imgb64 := base64.StdEncoding.EncodeToString([]byte(img))
		// replace the image url with base64 string
		markdown = strings.Replace(markdown, url[1], "data:image/png;base64,"+imgb64, 1)
	}
	return markdown
}

// mkdir
func Mkdir(path string) {
	os.MkdirAll(path, os.ModePerm)
}

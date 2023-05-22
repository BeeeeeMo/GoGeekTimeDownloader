package pkg

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
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
	regex := regexp.MustCompile(`!\[.*\]\((.*)\?.+\)`)
	// enumerate all the image urls
	for _, url := range regex.FindAllStringSubmatch(markdown, -1) {
		img := Fetch(url[1])
		extension, err := getExtensionFromUrl(url[1])
		if err != nil {
			log.Error(err)
			continue
		}
		imgb64 := base64.StdEncoding.EncodeToString([]byte(img))
		markdown = strings.Replace(markdown, url[1], "data:image/"+extension+";base64,"+imgb64, 1)
	}
	return markdown
}

func getExtensionFromUrl(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("could not parse URL: %v", err)
	}
	filename := path.Base(u.Path)
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return "", nil
	}
	extension := filename[index+1:]
	return extension, nil
}

// mkdir
func Mkdir(path string) {
	os.MkdirAll(path, os.ModePerm)
}

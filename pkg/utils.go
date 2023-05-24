package pkg

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/longbridgeapp/opencc"
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
	return markdown
}

func MdImgToLocal(markdown, path string) string {
	regex := regexp.MustCompile(`!\[.*\]\((.*)\)`)
	// enumerate all the image urls
	for _, url := range regex.FindAllStringSubmatch(markdown, -1) {
		img := Fetch(url[1])
		extension, err := getExtensionFromUrl(url[1])
		if err != nil {
			log.Error(err)
			continue
		}
		imgName := genStringMD5(url[1]) + "." + extension
		WriteToFile(path+"/images/"+imgName, img)
		markdown = strings.Replace(markdown, url[1], "../images/"+imgName, -1)
		log.Info(url[1])
	}
	return markdown
}

func genStringMD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
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
	// remove query string
	index = strings.Index(extension, "?")
	if index != -1 {
		extension = extension[:index]
	}
	return extension, nil
}

// mkdir
func Mkdir(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func S2T(s string) string {
	s2t, err := opencc.New("s2t")
	if err != nil {
		log.Fatal(err)
	}

	out, err := s2t.Convert(s)
	if err != nil {
		log.Fatal(err)
	}
	return out
}

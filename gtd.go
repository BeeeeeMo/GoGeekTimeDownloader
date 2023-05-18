package main

import (
	"encoding/json"
	"fmt"
	"gtd/pkg"
	"gtd/structs"
	"os"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	Headers  = map[string]string{}
	CourseId = 0
)

func init() {
	Headers["cookie"] = ""
	Headers["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
	CourseId = 0
}

func fetchArticle(articleId int, Headers map[string]string) string {
	Headers["Origin"] = "https://time.geekbang.org"
	Headers["Content-Type"] = "application/json"
	data := fmt.Sprintf(`{"id":"%v","include_neighbors":true,"is_freelyread":true}`, articleId)
	return pkg.PostWithHeaders("https://time.geekbang.org/serv/v1/article", data, Headers)
}

func fetchChapters(courseId int, Headers map[string]string) string {
	Headers["Origin"] = "https://time.geekbang.org"
	Headers["Content-Type"] = "application/json"
	data := fmt.Sprintf(`{"cid":"%v"}`, courseId)
	return pkg.PostWithHeaders("https://time.geekbang.org/serv/v1/chapters", data, Headers)
}

func fetchArticles(ChapterIds []string, Headers map[string]string) string {
	Headers["Origin"] = "https://time.geekbang.org"
	Headers["Content-Type"] = "application/json"
	data := fmt.Sprintf(`{"cid":48,"size":100,"prev":0,"order":"earliest","sample":false,"chapter_ids":["%v"]}`, strings.Join(ChapterIds, `","`))
	return pkg.PostWithHeaders("https://time.geekbang.org/serv/v1/column/articles", data, Headers)
}

func fetchCourseInfo(courseId int, Headers map[string]string) string {
	Headers["Origin"] = "https://time.geekbang.org"
	Headers["Content-Type"] = "application/json"
	data := fmt.Sprintf(`{"product_id":%v,"with_recommend_article":true}`, courseId)
	return pkg.PostWithHeaders("https://time.geekbang.org/serv/v3/column/info", data, Headers)
}

func main() {
	courseStr := fetchCourseInfo(CourseId, Headers)
	var course structs.CourseInfo

	json.Unmarshal([]byte(courseStr), &course)
	courseTitle := course.Data.Title
	pkg.Mkdir(courseTitle)

	fetchChaptersStr := fetchChapters(course.Data.Extra.Cid, Headers)
	var chapters structs.Chapters
	json.Unmarshal([]byte(fetchChaptersStr), &chapters)
	chapterIds := []string{}
	for _, chapter := range chapters.Data {
		fmt.Printf("chapter: %v\n", chapter.Title)
		chapterIds = append(chapterIds, chapter.ID)
	}

	fetchArticlesStr := fetchArticles(chapterIds, Headers)
	var articles structs.Articles
	json.Unmarshal([]byte(fetchArticlesStr), &articles)

	for _, article := range articles.Data.List {
		log.Info(article.ArticleTitle)
		articleStr := fetchArticle(article.ID, Headers)
		var article structs.Article
		json.Unmarshal([]byte(articleStr), &article)
		articleTitle := article.Data.ArticleTitle
		var subDir string
		for i, v := range chapters.Data {
			if v.ID == article.Data.ChapterID {
				subDir = fmt.Sprintf("%v-%v", i+1, v.Title)
				break
			}
		}
		fileName := courseTitle + "/" + subDir + "/" + articleTitle

		// check if file exists
		if _, err := os.Stat(fileName + ".mp3"); err == nil {
			log.Info("File exists, skip")
			time.Sleep(1 * time.Second)
			continue
		}

		markdown := pkg.HtmlToMarkdown(article.Data.ArticleContent)
		pkg.Mkdir(courseTitle + "/" + subDir)
		pkg.WriteToFile(fileName+".md", markdown)
		pkg.Download(fileName+".mp3", article.Data.AudioDownloadURL)
	}
}

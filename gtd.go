package main

var (
	Headers = map[string]string{}
)

func init() {
	Headers["cookie"] = ""
}

func fetchArticle(articleId int, Headers map[string]string) {
	Headers["Origin"] = "https://time.geekbang.org"
	Headers["Content-Type"] = "application/json"

}

func main() {

}

package pkg

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func DownloadFile(url string, path string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		log.Error(err)
		return
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Error(err)
		return
	}
}

func Fetch(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	log.Info("Fetching: ", url)
	if err != nil {
		log.Error(err)
		return ""
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return ""
	}

	if resp.StatusCode != 200 {
		log.Error(fmt.Sprintf("Status code error: %v", resp.StatusCode))
		return ""
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
		return ""
	}
	bodyString := string(bodyBytes)
	return bodyString
}

func FetchWithHeaders(url string, headers map[string]string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	log.Info("Fetching: ", url)
	if err != nil {
		log.Error(err)
		return ""
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return ""
	}

	if resp.StatusCode != 200 {
		log.Error(fmt.Sprintf("Status code error: %v", resp.StatusCode))
		return ""
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
		return ""
	}
	bodyString := string(bodyBytes)
	return bodyString
}

func Post(url string, data string) string {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	log.Info("Fetching: ", url)
	if err != nil {
		log.Error(err)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return ""
	}

	if resp.StatusCode != 200 {
		log.Error(fmt.Sprintf("Status code error: %v", resp.StatusCode))
		return ""
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
		return ""
	}
	bodyString := string(bodyBytes)
	return bodyString
}

func PostWithHeaders(url string, data string, headers map[string]string) string {
	client := &http.Client{}
	dataReader := io.Reader(strings.NewReader(data))
	req, err := http.NewRequest("POST", url, dataReader)
	log.Info("Fetching: ", url)
	if err != nil {
		log.Error(err)
		return ""
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return ""
	}

	if resp.StatusCode != 200 {
		log.Error(fmt.Sprintf("Status code error: %v", resp.StatusCode))
		return ""
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Error(err)
		return ""
	}
	bodyString := string(bodyBytes)
	return bodyString
}

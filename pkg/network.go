package pkg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetReportCaller(true)
	log.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", f.File, f.Line)
		},
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

func retryRequests(req *http.Request, retry int) (*http.Response, error) {
	client := &http.Client{}
	var resp *http.Response
	var err error

	for i := 0; i <= retry; i++ {
		resp, err = client.Do(req)

		if err == nil && resp.StatusCode == http.StatusOK {
			return resp, nil
		}

		log.Error(fmt.Sprintf("Request failure: %v, retrying (%d/%d)..", err, i+1, retry+1))
		time.Sleep(time.Second * 10)
	}

	return resp, errors.New(fmt.Sprintf("Max retries reached (%d). Last error: %v", retry, err))
}

func FetchWithHeaders(url string, headers map[string]string, retry int) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return ""
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := retryRequests(req, retry)

	if err != nil {
		log.Error(err)
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

func PostWithHeaders(url string, data string, headers map[string]string, retry int) string {
	dataReader := io.Reader(strings.NewReader(data))
	req, err := http.NewRequest("POST", url, dataReader)
	if err != nil {
		log.Error(err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := retryRequests(req, retry)

	if err != nil {
		log.Error(err, data)
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

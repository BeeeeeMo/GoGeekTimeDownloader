package pkg

import (
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

func DownloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	progress := progressbar.NewOptions(
		size,
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetDescription("Downloading:"),
		progressbar.OptionSetTheme(progressbar.Theme{Saucer: "=", SaucerHead: ">", SaucerPadding: " ", BarStart: "|", BarEnd: "|"}),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetPredictTime(true),
	)

	_, err = io.Copy(io.MultiWriter(file, progress), resp.Body)
	return err
}

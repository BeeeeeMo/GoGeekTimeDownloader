package pkg

import "regexp"

func CleanFileName(name string) string {
	regex := regexp.MustCompile(`[\\/:*?"<>|]`)
	return regex.ReplaceAllString(name, "")
}

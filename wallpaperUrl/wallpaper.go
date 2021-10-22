package wallpaperUrl

import (
	"strings"
)

const UNSPLASH_SOURCE = "https://source.unsplash.com/"

func GetRandomWallpaperURL(resolution string, keywords []string) string {
	sourceUrl := strings.Join([]string{UNSPLASH_SOURCE, "random/", resolution, "/"}, "")

	if len(keywords) > 0 {
		sourceUrl = sourceUrl + "?" + strings.Join(keywords, ",")
	}

	return sourceUrl
}

func GetDailyWallpaperURL(resolution string, keywords []string) string {
	sourceUrl := strings.Join([]string{UNSPLASH_SOURCE, resolution, "/daily/"}, "")

	if len(keywords) > 0 {
		sourceUrl = sourceUrl + "?" + strings.Join(keywords, ",")
	}

	return sourceUrl
}

func GetWeeklyWallpaperURL(resolution string, keywords []string) string {
	sourceUrl := strings.Join([]string{UNSPLASH_SOURCE, resolution, "/weekly/"}, "")

	if len(keywords) > 0 {
		sourceUrl = sourceUrl + "?" + strings.Join(keywords, ",")
	}

	return sourceUrl
}

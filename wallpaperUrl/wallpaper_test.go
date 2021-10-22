package wallpaperUrl

import "testing"

const typicalResolution string = "1920x1080"

func TestRandomWallpaperUrl(t *testing.T) {
	url := GetRandomWallpaperURL(typicalResolution, []string{})

	if url != "https://source.unsplash.com/random/1920x1080/" {
		t.Error("Malformed URL", url)
	}
}

func TestRandomWallpaperUrlKeywords(t *testing.T) {
	url := GetRandomWallpaperURL(typicalResolution, []string{"lorem", "ipsum"})

	if url != "https://source.unsplash.com/random/1920x1080/?lorem,ipsum" {
		t.Error("Malformed URL", url)
	}
}

func TestDailyWallpaperUrl(t *testing.T) {
	url := GetDailyWallpaperURL(typicalResolution, []string{})

	if url != "https://source.unsplash.com/1920x1080/daily/" {
		t.Error("Malformed URL", url)
	}
}

func TestDailyWallpaperUrlKeywords(t *testing.T) {
	url := GetDailyWallpaperURL(typicalResolution, []string{"lorem", "ipsum"})

	if url != "https://source.unsplash.com/1920x1080/daily/?lorem,ipsum" {
		t.Error("Malformed URL", url)
	}
}

func TestWeeklyWallpaperUrl(t *testing.T) {
	url := GetWeeklyWallpaperURL(typicalResolution, []string{})

	if url != "https://source.unsplash.com/1920x1080/weekly/" {
		t.Error("Malformed URL", url)
	}
}

func TestWeeklyWallpaperUrlKeywords(t *testing.T) {
	url := GetWeeklyWallpaperURL(typicalResolution, []string{"lorem", "ipsum"})

	if url != "https://source.unsplash.com/1920x1080/weekly/?lorem,ipsum" {
		t.Error("Malformed URL", url)
	}
}

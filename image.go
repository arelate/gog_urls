package gog_urls

import (
	"net/url"
	"strings"
)

const (
	pngExt               = ".png"
	jpgExt               = ".jpg"
	formatterPlaceholder = "_{formatter}"
)

func Image(srcImage string) (*url.URL, error) {
	imgUrl, err := url.Parse(srcImage)
	if err != nil {
		return nil, err
	}

	imgUrl.Scheme = HttpsScheme

	if strings.Contains(imgUrl.Path, formatterPlaceholder) {
		imgUrl.Path = strings.Replace(imgUrl.Path, formatterPlaceholder, "", 1)
	}

	// make sure we're always downloading .png source image
	if strings.HasSuffix(imgUrl.Path, jpgExt) {
		imgUrl.Path = strings.Replace(imgUrl.Path, jpgExt, pngExt, 1)
	}

	if !strings.HasSuffix(imgUrl.Path, pngExt) {
		imgUrl.Path += pngExt
	}

	return imgUrl, nil
}

func Screenshots(csvScreenshots string) ([]*url.URL, error) {
	screenshots := strings.Split(csvScreenshots, ",")
	scrUrls := make([]*url.URL, 0, len(screenshots))
	for _, scr := range screenshots {
		scrUrl, err := Image(scr)
		if err != nil {
			return scrUrls, err
		}
		scrUrls = append(scrUrls, scrUrl)
	}
	return scrUrls, nil
}

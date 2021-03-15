package gog_urls

import (
	"net/url"
	"os"
	"path"
	"strings"
)

const (
	pngExt               = ".png"
	jpgExt               = ".jpg"
	formatterPlaceholder = "_{formatter}"
)

func ImageId(srcImage string) string {
	if srcImage == "" {
		return srcImage
	}
	var lastPart string
	pathParts := strings.Split(srcImage, (string)(os.PathSeparator))
	if len(pathParts) > 0 {
		lastPart = pathParts[len(pathParts)-1]
	}

	return strings.TrimSuffix(lastPart, path.Ext(lastPart))
}

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

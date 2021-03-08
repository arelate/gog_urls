package gog_urls

import (
	"net/url"
	"strings"
)

const (
	pngExt = ".png"
	jpgExt = ".jpg"
)

func Image(srcImage string) (*url.URL, error) {
	imgUrl, err := url.Parse(srcImage)
	if err != nil {
		return nil, err
	}

	imgUrl.Scheme = HttpsScheme

	// make sure we're always downloading .png source image
	if strings.HasSuffix(imgUrl.Path, jpgExt) {
		imgUrl.Path = strings.Replace(imgUrl.Path, jpgExt, pngExt, 1)
	}

	if !strings.HasSuffix(imgUrl.Path, pngExt) {
		imgUrl.Path += pngExt
	}

	return imgUrl, nil
}

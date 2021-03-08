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

	if !strings.HasSuffix(imgUrl.Path, pngExt) &&
		!strings.HasSuffix(imgUrl.Path, jpgExt) {
		imgUrl.Path += pngExt
	}

	return imgUrl, nil
}

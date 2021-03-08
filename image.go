package gog_urls

import (
	"net/url"
	"strings"
)

const pngExt = ".png"

func Image(srcImage string) (*url.URL, error) {
	imgUrl, err := url.Parse(srcImage)
	if err != nil {
		return nil, err
	}

	imgUrl.Scheme = HttpsScheme

	if !strings.HasSuffix(imgUrl.Path, pngExt) {
		imgUrl.Path += pngExt
	}

	return imgUrl, nil
}

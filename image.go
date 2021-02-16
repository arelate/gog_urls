package gog_urls

import (
	"net/url"
)

const pngExt = ".png"

func Image(srcImage string) (*url.URL, error) {
	imgUrl, err := url.Parse(srcImage)
	if err != nil {
		return nil, err
	}

	imgUrl.Scheme = HttpsScheme
	imgUrl.Path += pngExt

	return imgUrl, nil
}

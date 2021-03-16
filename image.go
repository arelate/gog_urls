package gog_urls

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

const (
	formatterPlaceholder = "_{formatter}"
	pngExt               = ".png"
	jpgExt               = ".jpg"
)

func ImageId(srcImage string) string {
	if srcImage == "" {
		return srcImage
	}

	_, fn := path.Split(srcImage)
	fnSansExt := strings.TrimSuffix(fn, path.Ext(fn))

	return strings.TrimSuffix(fnSansExt, formatterPlaceholder)
}

func Image(imageId string) (*url.URL, error) {
	if imageId == "" {
		return nil, fmt.Errorf("gog_urls: empty image-id")
	}
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   imagesHost,
		Path:   strings.Replace(imagesPathTemplate, "{0}", imageId, 1),
	}, nil
}

func Screenshot(imageId string) (*url.URL, error) {
	imgUrl, err := Image(imageId)
	if err != nil {
		return imgUrl, err
	}
	imgUrl.Path = strings.Replace(imgUrl.Path, pngExt, jpgExt, 1)
	return imgUrl, err
}

func Screenshots(screenshots []string) ([]*url.URL, error) {
	scrUrls := make([]*url.URL, 0, len(screenshots))
	for _, scr := range screenshots {
		scrUrl, err := Screenshot(scr)
		if err != nil {
			return scrUrls, err
		}
		scrUrls = append(scrUrls, scrUrl)
	}
	return scrUrls, nil
}

package gog_urls

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

const (
	pngExt               = ".png"
	formatterPlaceholder = "_{formatter}"
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
		Path:   imageId + pngExt,
	}, nil
}

//func Image(srcImage string) (*url.URL, error) {
//	imgUrl, err := url.Parse(srcImage)
//	if err != nil {
//		return nil, err
//	}
//
//	imgUrl.Scheme = HttpsScheme
//
//	if strings.Contains(imgUrl.Path, formatterPlaceholder) {
//		imgUrl.Path = strings.Replace(imgUrl.Path, formatterPlaceholder, "", 1)
//	}
//
//	// make sure we're always downloading .png source image
//	if strings.HasSuffix(imgUrl.Path, jpgExt) {
//		imgUrl.Path = strings.Replace(imgUrl.Path, jpgExt, pngExt, 1)
//	}
//
//	if !strings.HasSuffix(imgUrl.Path, pngExt) {
//		imgUrl.Path += pngExt
//	}
//
//	return imgUrl, nil
//}

func Screenshots(screenshots []string) ([]*url.URL, error) {
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

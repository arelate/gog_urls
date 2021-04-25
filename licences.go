package gog_urls

import (
	"github.com/arelate/gog_media"
	"net/url"
)

func Licences() *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   menuHost,
		Path:   licencesPath,
	}
}

func DefaultLicences(id string, mt gog_media.Media) *url.URL {
	return Licences()
}

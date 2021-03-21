package gog_urls

import (
	"github.com/arelate/gog_media"
	"net/url"
)

var expandValues = []string{
	"downloads",
	"expanded_dlcs",
	"description",
	"screenshots",
	"videos",
	"related_products",
	"changelog"}

func ApiProductV2(id string, mt gog_media.Media) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   apiHost,
		Path:   apiV2GamesPath + id,
	}
}

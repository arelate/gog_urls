package gog_urls

import (
	"github.com/arelate/gog_types"
	"net/url"
)

func ApiProductV2(id string, mt gog_types.Media) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   apiHost,
		Path:   apiV2GamesPath + id,
	}
}

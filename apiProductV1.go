package gog_urls

import (
	"github.com/arelate/gog_types"
	"net/url"
	"strings"
)

func ApiProductV1(id string, mt gog_types.Media) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   apiHost,
		Path:   strings.Replace(apiV1PathTemplate, "{id}", id, 1),
	}
}

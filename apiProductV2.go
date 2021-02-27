package gog_urls

import (
	"github.com/arelate/gog_types"
	"net/url"
	"strings"
)

var expandValues = []string{
	"downloads",
	"expanded_dlcs",
	"description",
	"screenshots",
	"videos",
	"related_products",
	"changelog"}

func ApiProductV2(id string, mt gog_types.Media) *url.URL {
	apv2url := &url.URL{
		Scheme: HttpsScheme,
		Host:   apiHost,
		Path:   apiV2GamesPath + id,
	}

	q := apv2url.Query()
	q.Set("expand", strings.Join(expandValues, ","))
	apv2url.RawQuery = q.Encode()

	return apv2url
}

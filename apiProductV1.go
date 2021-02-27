package gog_urls

import (
	"github.com/arelate/gog_types"
	"net/url"
	"strings"
)

func ApiProductV1(id string, mt gog_types.Media) *url.URL {
	apv1url := &url.URL{
		Scheme: HttpsScheme,
		Host:   apiHost,
		Path:   strings.Replace(apiV1PathTemplate, "{id}", id, 1),
	}

	q := apv1url.Query()
	q.Set("expand", strings.Join(expandValues, ","))
	apv1url.RawQuery = q.Encode()

	return apv1url
}

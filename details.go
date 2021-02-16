// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_urls

import (
	"github.com/arelate/gog_types"
	"net/url"
	"strings"
)

func Details(id string, mt gog_types.Media) *url.URL {
	path := strings.Replace(detailsPathTemplate, "{mediaType}", mt.String(), 1)
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(path, "{id}", id, 1),
	}
}

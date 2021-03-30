// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_urls

import (
	"github.com/arelate/gog_media"
	"github.com/arelate/gog_types"
	"net/url"
)

func DefaultStorePage(
	page string,
	mt gog_media.Media) *url.URL {
	return StorePage(page, mt, gog_types.StoreSortByNewestFirst)
}

func StorePage(
	page string,
	mt gog_media.Media,
	sortOrder gog_types.StoreSortOrder) *url.URL {

	storePage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   storePagePath,
	}

	q := storePage.Query()
	q.Add("mediaType", mt.String())
	if sortOrder != "" {
		q.Add("sort", string(sortOrder))
	}
	q.Add("page", page)
	storePage.RawQuery = q.Encode()

	return storePage
}

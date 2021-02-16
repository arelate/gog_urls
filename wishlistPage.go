// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_urls

import (
	"github.com/arelate/gog_types"
	"net/url"
	"strconv"
)

func DefaultWishlistPage(
	page string,
	mt gog_types.Media) *url.URL {
	return WishlistPage(page, mt, gog_types.WishlistSortByDateAdded, false)
}

func WishlistPage(
	page string,
	mt gog_types.Media,
	sortOrder gog_types.WishlistSortOrder,
	hidden bool) *url.URL {

	wishlistPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   wishlistPath,
	}

	q := wishlistPage.Query()
	q.Add("mediaType", strconv.Itoa(int(mt)))
	if sortOrder != "" {
		q.Add("sortBy", string(sortOrder))
	}
	if hidden {
		q.Add("hiddenFlag", "1")
	}
	q.Add("page", page)
	wishlistPage.RawQuery = q.Encode()

	return wishlistPage
}

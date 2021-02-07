// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogurls

import (
	"github.com/arelate/gogtypes"
	"net/url"
	"strconv"
)

func DefaultWishlistPage(
	page string,
	mt gogtypes.Media) *url.URL {
	return WishlistPage(page, mt, gogtypes.WishlistSortByDateAdded, false)
}

func WishlistPage(
	page string,
	mt gogtypes.Media,
	sortOrder gogtypes.WishlistSortOrder,
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

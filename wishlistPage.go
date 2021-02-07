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
	page int,
	mt gogtypes.Media) *url.URL {
	return WishlistPage(page, mt, gogtypes.WishlistSortByDateAdded, false)
}

func WishlistPage(
	page int,
	mt gogtypes.Media,
	sortOrder gogtypes.WishlistSortOrder,
	hidden bool) *url.URL {

	wishlistPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   wishlistPath,
	}

	hiddenFlag := "0"
	if hidden {
		hiddenFlag = "1"
	}

	if sortOrder == "" {
		sortOrder = gogtypes.WishlistSortByDateAdded
	}

	q := wishlistPage.Query()
	q.Add("mediaType", strconv.Itoa(int(mt)))
	q.Add("sortBy", string(sortOrder))
	q.Add("hiddenFlag", hiddenFlag)
	q.Add("page", strconv.Itoa(page))
	wishlistPage.RawQuery = q.Encode()

	return wishlistPage
}

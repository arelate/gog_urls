// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_urls

import (
	"github.com/arelate/gog_media"
	"github.com/arelate/gog_types"
	"net/url"
	"strconv"
	"strings"
)

func DefaultWishlistPage(
	page string,
	mt gog_media.Media) *url.URL {
	return WishlistPage(page, mt, gog_types.WishlistSortByDateAdded, false)
}

func WishlistPage(
	page string,
	mt gog_media.Media,
	sortOrder gog_types.WishlistSortOrder,
	hidden bool) *url.URL {

	wishlistPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   wishlistSearchPath,
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

func AddToWishlist(id string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(wishlistAddPathTemplate, "{id}", id, 1),
	}
}

func RemoveFromWishlist(id string) *url.URL {
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   strings.Replace(wishlistRemovePathTemplate, "{id}", id, 1),
	}
}

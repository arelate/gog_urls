// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_urls

import (
	"github.com/arelate/gog_types"
	"net/url"
	"strconv"
)

func DefaultAccountProductsPage(
	page string,
	mt gog_types.Media) *url.URL {
	return AccountProductsPage(
		page,
		mt,
		gog_types.AccountProductsSortPurchaseDate,
		false,
		false)
}

func AccountProductsPage(
	page string,
	mt gog_types.Media,
	sortOrder gog_types.AccountProductsSortOrder,
	updated bool, /* get only updated products */
	hidden bool /* get only hidden products */) *url.URL {

	accountProductsPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   accountProductsPagePath,
	}

	q := accountProductsPage.Query()
	q.Add("mediaType", strconv.Itoa(int(mt)))
	if sortOrder != "" {
		q.Add("sortBy", string(sortOrder))
	}
	q.Add("page", page)
	if hidden {
		q.Add("hiddenFlag", "1")
	}
	if updated {
		q.Add("isUpdated", "1")
	}
	accountProductsPage.RawQuery = q.Encode()

	return accountProductsPage
}

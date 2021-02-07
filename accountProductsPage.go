// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogurls

import (
	"github.com/arelate/gogtypes"
	"net/url"
	"strconv"
)

func DefaultAccountProductsPage(
	page string,
	mt gogtypes.Media) *url.URL {
	return AccountProductsPage(
		page,
		mt,
		gogtypes.AccountProductsSortPurchaseDate,
		false,
		false)
}

func AccountProductsPage(
	page string,
	mt gogtypes.Media,
	sortOrder gogtypes.AccountProductsSortOrder,
	updated bool, /* get only updated products */
	hidden bool /* get only hidden products */) *url.URL {

	accountProductsPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   accountProductsPagePath,
	}

	if sortOrder == "" {
		sortOrder = gogtypes.AccountProductsSortPurchaseDate
	}

	q := accountProductsPage.Query()
	q.Add("mediaType", strconv.Itoa(int(mt)))
	q.Add("sortBy", string(sortOrder))
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

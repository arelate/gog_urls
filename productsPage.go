// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_urls

import (
	"github.com/arelate/gog_types"
	"net/url"
)

func DefaultProductsPage(
	page string,
	mt gog_types.Media) *url.URL {
	return ProductsPage(page, mt, gog_types.ProductsSortByNewestFirst)
}

func ProductsPage(
	page string,
	mt gog_types.Media,
	sortOrder gog_types.ProductsSortOrder) *url.URL {

	productsPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   productsPagePath,
	}

	q := productsPage.Query()
	q.Add("mediaType", mt.String())
	if sortOrder != "" {
		q.Add("sort", string(sortOrder))
	}
	q.Add("page", page)
	productsPage.RawQuery = q.Encode()

	return productsPage
}

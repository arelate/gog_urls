// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogurls

import (
	"github.com/arelate/gogtypes"
	"net/url"
)

func DefaultProductsPage(
	page string,
	mt gogtypes.Media) *url.URL {
	return ProductsPage(page, mt, gogtypes.ProductsSortByNewestFirst)
}

func ProductsPage(
	page string,
	mt gogtypes.Media,
	sortOrder gogtypes.ProductsSortOrder) *url.URL {

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

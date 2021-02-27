// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_urls

const (
	HttpsScheme = "https"
)

// hosts
const (
	GogHost    = "gog.com"
	WwwGogHost = "www." + GogHost
	apiHost    = "api.gog.com"
)

// paths
const (
	detailsPathTemplate     = "/account/{mediaType}Details/{id}.json"
	productsPagePath        = "/games/ajax/filtered"
	accountProductsPagePath = "/account/getFilteredProducts"
	wishlistPath            = "/account/wishlist/search"
	apiV1PathTemplate       = "/products/{id}?expand=downloads,expanded_dlcs,description,screenshots,videos,related_products,changelog"
	apiV2GamesPath          = "/v2/games/"
)

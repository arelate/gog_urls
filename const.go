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
	imagesHost = "images.gog-statics.com"
	menuHost   = "menu.gog.com"
)

// paths
const (
	storePagePath              = "/games/ajax/filtered"
	accountPath                = "/account"
	detailsPathTemplate        = accountPath + "/{mediaType}Details/{id}.json"
	accountPagePath            = accountPath + "/getFilteredProducts"
	accountWishlistPath        = accountPath + "/wishlist"
	wishlistSearchPath         = accountWishlistPath + "/search"
	userWishlistPath           = "/user/wishlist"
	wishlistAddPathTemplate    = userWishlistPath + "/add/{id}"
	wishlistRemovePathTemplate = userWishlistPath + "/remove/{id}"
	apiV1PathTemplate          = "/products/{id}"
	apiV2GamesPath             = "/v2/games/"
	imagesPathTemplate         = "/{image_id}.png"
	licencesPath               = "/v1/account/licences"
	accountTagsPath            = "/account/tags"
	tagsCreatePath             = accountTagsPath + "/add"
	tagsDeletePath             = accountTagsPath + "/delete"
	tagsUpdatePath             = accountTagsPath + "/update"
	tagsAddPath                = accountTagsPath + "/attach"
	tagsRemovePath             = accountTagsPath + "/detach"
)

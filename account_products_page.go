// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gog_urls

import (
	"github.com/arelate/gog_media"
	"github.com/arelate/gog_types"
	"net/url"
	"strconv"
)

func DefaultAccountPage(
	page string,
	mt gog_media.Media) *url.URL {
	return AccountPage(
		page,
		mt,
		gog_types.AccountSortByPurchaseDate,
		false,
		false)
}

func AccountPage(
	page string,
	mt gog_media.Media,
	sortOrder gog_types.AccountSortOrder,
	updated bool, /* get only updated products */
	hidden bool /* get only hidden products */) *url.URL {

	accountPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   accountPagePath,
	}

	q := accountPage.Query()
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
	accountPage.RawQuery = q.Encode()

	return accountPage
}

// Copyright (c) 2020. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be
// found in the LICENSE file.

package gogurls

import (
	"fmt"
	"github.com/arelate/gogtypes"
	"net/url"
	"strings"
)

func Details(id int, mt gogtypes.Media) *url.URL {
	path := strings.Replace(detailsPath, "{mediaType}", mt.String(), 1)
	return &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   fmt.Sprintf("%s%d.json", path, id),
	}
}

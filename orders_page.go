package gog_urls

import (
	"github.com/arelate/gog_media"
	"net/url"
)

const (
	pageParam        = "page"
	canceledParam    = "canceled"
	completedParam   = "completed"
	inProgressParam  = "in_progress"
	notRedeemedParam = "not_redeemed"
	pendingParam     = "pending"
	redeemedParam    = "redeemed"
)

func boolToIntStr(flag bool) string {
	if flag {
		return "1"
	}
	return "0"
}

func DefaultOrdersPage(
	page string,
	mt gog_media.Media) *url.URL {
	return OrdersPage(
		page,
		true,
		true,
		true,
		true,
		true,
		true)
}

func OrdersPage(page string, canceled, completed, inProgress, notRedeemed, pending, redeemed bool) *url.URL {
	orderPage := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   ordersPath,
	}

	q := orderPage.Query()
	q.Add(pageParam, page)
	q.Add(canceledParam, boolToIntStr(canceled))
	q.Add(completedParam, boolToIntStr(completed))
	q.Add(inProgressParam, boolToIntStr(inProgress))
	q.Add(notRedeemedParam, boolToIntStr(notRedeemed))
	q.Add(pendingParam, boolToIntStr(pending))
	q.Add(redeemedParam, boolToIntStr(redeemed))
	orderPage.RawQuery = q.Encode()

	return orderPage
}

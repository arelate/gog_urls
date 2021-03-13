package gog_urls

import "net/http"

const (
	AcceptTextHTML = "text/html"
	AcceptImages   = "image/webp,image/png,image/svg+xml,image/*"
)

func SetDefaultHeaders(req *http.Request, host, accept string) {
	const (
		acceptLanguageHeader = "en-us"
		connectionHeader     = "keep-alive"
		userAgentHeader      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_3_0) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/89.0.4389.82 " +
			"Safari/537.36 " +
			"Edg/89.0.774.50" // Microsoft Edge 89 UA string
	)
	if host != "" {
		req.Host = host
	}
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	req.Header.Set("Accept-Language", acceptLanguageHeader)
	req.Header.Set("Connection", connectionHeader)
	req.Header.Set("User-Agent", userAgentHeader)
	req.Header.Set("Origin", WwwGogHost)
}

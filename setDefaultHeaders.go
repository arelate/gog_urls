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
		userAgentHeader      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/87.0.4280.141 " +
			"Safari/537.36 " +
			"Edg/87.0.664.75" // Microsoft Edge 87 UA string
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

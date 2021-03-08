package gog_urls

import "net/url"

func BoxArt(srcBoxArt string) (*url.URL, error) {
	return url.Parse(srcBoxArt)
}

module github.com/arelate/gog_urls

go 1.16

require (
	github.com/arelate/gog_media v0.1.2-alpha
	github.com/arelate/gog_types v0.1.6-alpha
)

replace (
	github.com/arelate/gog_media => ../gog_media
	github.com/arelate/gog_types => ../gog_types

)

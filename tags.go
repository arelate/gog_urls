package gog_urls

import "net/url"

func CreateTag(name string) *url.URL {
	createTag := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   tagsCreatePath,
	}

	q := createTag.Query()
	q.Add("name", name)
	createTag.RawQuery = q.Encode()

	return createTag
}

func DeleteTag(tagId string) *url.URL {
	deleteTag := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   tagsDeletePath,
	}

	q := deleteTag.Query()
	q.Add("tag_id", tagId)
	deleteTag.RawQuery = q.Encode()

	return deleteTag
}

func UpdateTags(tagsJson string) *url.URL {
	updateTags := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   tagsUpdatePath,
	}

	q := updateTags.Query()
	q.Add("tags", tagsJson)
	updateTags.RawQuery = q.Encode()

	return updateTags
}

func AddTag(productId, tagId string) *url.URL {
	addTag := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   tagsAddPath,
	}

	q := addTag.Query()
	q.Add("product_id", productId)
	q.Add("tag_id", tagId)
	addTag.RawQuery = q.Encode()

	return addTag
}

func RemoveTag(productId, tagId string) *url.URL {
	removeTag := &url.URL{
		Scheme: HttpsScheme,
		Host:   GogHost,
		Path:   tagsRemovePath,
	}

	q := removeTag.Query()
	q.Add("product_id", productId)
	q.Add("tag_id", tagId)
	removeTag.RawQuery = q.Encode()

	return removeTag
}

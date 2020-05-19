package storage

type StorageImage interface {
	GetBaseImageUrl() string
	SetFullImageUrl(string2 string)
}

func AutoComplementImageUrl(image StorageImage) {
	url := _bucket.AccessUrl() + image.GetBaseImageUrl()
	image.SetFullImageUrl(url)
}

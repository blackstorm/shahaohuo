package storage

type StorageImage interface {
	GetBaseImageUrl() string
}

func ComplementImageUrl(image StorageImage) string {
	return _bucket.AccessUrl() + image.GetBaseImageUrl()
}

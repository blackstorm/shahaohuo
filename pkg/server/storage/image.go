package storage

const StorageEndPoint = "http://118.24.52.45:30434/shahaohuo"

type StorageImage interface {
	GetBaseImageUrl() string
}

func ComplementImageUrl(image StorageImage) string {
	return StorageEndPoint + image.GetBaseImageUrl()
}

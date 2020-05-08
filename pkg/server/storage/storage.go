package storage

import "shahaohuo.com/shahaohuo/pkg/bucket"

var _bucket *bucket.Bucket

func InitStorage(bucket *bucket.Bucket) {
	_bucket = bucket
}

func GetBucket() *bucket.Bucket {
	return _bucket
}

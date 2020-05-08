package bucket

import "errors"

const (
	BucketNotExist        = "BucketNotExist"
	CreateBucketFailed    = "CreateBucketFailed"
	PutFileToBucketFailed = "PutFileToBucketFailed"
)

var BucketNotExistError = errors.New(BucketNotExist)
var CreateBucketFailedError = errors.New(CreateBucketFailed)
var PutFileToBucketFailedError = errors.New(PutFileToBucketFailed)

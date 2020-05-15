package util

type AsyncResult struct {
	Ret   interface{}
	Error error
}

func CheckAsyncResultsError(res ...AsyncResult) bool {
	for _, ret := range res {
		if ret.Error != nil {
			return true
		}
	}
	return false
}

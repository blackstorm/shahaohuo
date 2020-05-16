package image

import "shahaohuo.com/shahaohuo/pkg/server/orm"

func CheckIsUserImage(userId, path string) bool {
	if image, e := orm.FindByUserIdAndPath(userId, path); e != nil {
		return false
	} else {
		return image != nil
	}
}

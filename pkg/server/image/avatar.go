package image

import (
	"strconv"
	"time"
)

func UserAvatarPathGen(userId string) string {
	tz := time.Now().Unix()
	return "/users/" + userId + "/avatars/" + strconv.FormatInt(tz, 10)
}

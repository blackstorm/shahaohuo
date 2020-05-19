package image

import (
	"github.com/rs/xid"
)

func HaohuoImagePathRandom() string {
	id := xid.New().String()
	return "/haohuos/images/" + id
}

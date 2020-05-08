package seo

import "shahaohuo.com/shahaohuo/pkg/server/orm"

type KeyWorld interface {
	SeoKeyWorld() string
}

func BaseKeyWorlds(base string, tags []orm.Tag) string {
	ret := "啥好货," + base
	if tags != nil && len(tags) > 0 {
		for _, v := range tags {
			ret += "," + v.SeoKeyWorld()
		}
	}
	return ret
}

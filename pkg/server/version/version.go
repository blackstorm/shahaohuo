package version

import "strings"

type V struct {
	version string
}

var Version = &V{}

func SetVersion(version string) {
	Version.version = version
}

func (v *V) GetVersion() string {
	return v.version
}

func (v *V) StaticVersion(href string) string {
	if strings.Contains(href, "?") {
		href += "&v=" + v.version
	} else {
		href += "?v=" + v.version
	}
	return href
}

package util

type Mime string

func (m Mime) StringValue() string {
	return string(m)
}

const (
	PNG  Mime = "image/png"
	JPEG Mime = "image/jpeg"
)

func CheckIsContain(value string, types ...Mime) bool {
	for _, t := range types {
		if value == t.StringValue() {
			return true
		}
	}
	return false
}

package token

import (
	"github.com/dgrijalva/jwt-go"
)

var signSecret = []byte("uVQ5$JrUp%h*ZN*%CQ%GOgQsZ^HUADO!")

type Claims struct {
	jwt.StandardClaims
	Id string `json:"id"`
}

func Sign(id string, expiredAt int64) (string, error) {

	claims := &Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expiredAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(signSecret)
}

func Verify(str string) (string, bool) {
	if token, e := jwt.ParseWithClaims(str, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return signSecret, nil
	}); e == nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims.Id, true
		}
	}
	return "", false
}

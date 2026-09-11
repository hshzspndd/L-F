package utiles

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId   int
	UserName string
	jwt.RegisteredClaims
}

func GenerateJwt(userId int, userName string) (string, error, time.Time) {
	var (
		key                 = []byte("zjh20080210")
		expiredAt time.Time = time.Now().Add(2 * time.Hour)
	)
	claims := Claims{
		UserId:   userId,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
			Issuer:    "zjh",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err, time.Now()
	}

	return tokenString, nil, expiredAt

}

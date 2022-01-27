package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Id        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(Id uint, username string, authority int) (string, error) {
	notTime := time.Now()
	expireTime := notTime.Add(24 * time.Hour)
	//fmt.Println(notTime)
	//fmt.Println(expireTime)
	claims := Claims{
		Id:        Id,
		UserName:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "todo_list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		//fmt.Println("token: ",tokenClaims)
		//fmt.Println("claims: ",claims)
		//fmt.Println(ok)
		//fmt.Println(tokenClaims.Valid)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
		//fmt.Println("qqqqqqqqqqqqqq")
	}
	return nil, err
}

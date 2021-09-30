package jwt

import (
	"time"

	"github.com/SparkElf/blog-server/config"
	. "github.com/SparkElf/blog-server/utils/exception"
	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte(config.JwtKey)

//jwt token中包含的核验内容
type Claim struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	jwt.StandardClaims
}

func SetToken(username string, password string) string {
	expireTime := time.Now().Add(time.Duration(config.ExpireTime))
	//核验内容
	claims := Claim{
		username, password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "SparkElf", //签发者
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(JwtKey)
	return tokenString
}
func CheckToken(token string) (*Claim, Err) {
	t, _ := jwt.ParseWithClaims(token, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil //私钥加密 公钥解密 对HS256公钥私钥相同 对RS256公钥私钥不同并且可能有多个不同的公钥
	})
	claims, _ := t.Claims.(*Claim) //断言 类型转换 多态 Claims转MyClaims指针
	if t.Valid {
		if time.Now().Unix() > claims.ExpiresAt {
			return nil, Err{Code: ERROR_TOKEN_OVERDUE}
		}
		return claims, Err{Code: SUCCESE}
	} else {
		return nil, Err{Code: ERROR_TOKEN_WRONG}
	}
}

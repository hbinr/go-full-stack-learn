package jwt

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

const (
	// TokenExpireDuration  access token 过期时间
	TokenExpireDuration = time.Minute * 10
	// RefreshTokenExpireDuration 刷新 access token 过期时间
	RefreshTokenExpireDuration = time.Hour * 24 * 7
	// MySecret 私人秘钥
	MySecret = "full-stack-study"
)

type MyClaims struct {
	UserID int64 `json:"userID"`
	jwt.StandardClaims
}

// GenToken 生成access token和refresh token
func GenToken(userID int64) (aToken, rToken string, err error) {
	// 创建一个我们自己的声明
	c := &MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // aToken过期时间
			Issuer:    "XS-bbs",                                   // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	if aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).
		SignedString([]byte(MySecret)); err != nil { // 使用指定的secret签名，注意转换为字节切片，并获得完整的编码后的字符串token
		return "", "", err
	}

	// refresh token生成
	if rToken, err = jwt.NewWithClaims(jwt.SigningMethodES256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), // rToken过期时间
		Issuer:    "XS-bbs",
	}).SignedString([]byte(MySecret)); err != nil {
		return "", "", err
	}

	return aToken, rToken, err
}

// ParseToken 解析access token
func ParseToken(aToken string) (*MyClaims, error) {
	// 解析token
	var mc = new(MyClaims) // 这行代码一定要写，否则无法分配内存地址，在赋值的时候找不到值
	token, err := jwt.ParseWithClaims(aToken, mc, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

// RefreshToken 刷新access token
func RefreshToken(oldAToken, rToken string) (newAtoken string, err error) {
	// 1.refresh token无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		return
	}
	// 2.从旧access token中解析处cliams数据
	var claims MyClaims
	if _, err = jwt.ParseWithClaims(oldAToken, &claims, keyFunc); err != nil {
		v, ok := err.(*jwt.ValidationError)
		if ok {
			return "", errors.New("invalid token")
		}

		// 3.当旧access token过期，返回过期错误，并且refresh token没有过期就创建一个新的access token
		if v.Errors == jwt.ValidationErrorExpired {
			if newAtoken, _, err = GenToken(claims.UserID); err != nil {
				return "", err
			}
		}
	}
	return newAtoken, nil
}

func keyFunc(token *jwt.Token) (i interface{}, err error) {
	return MySecret, nil
}

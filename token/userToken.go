package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/geiqin/supports/auth"
	"time"
)

type UserAble interface {
	Decode(tokenStr string) (*UserClaims, error)
	Encode(user *auth.LoginUser) (string, error)
}


// 自定义的 metadata，在加密后作为 JWT 的第二部分返回给客户端
type UserClaims struct {
	User *auth.LoginUser
	Limit *auth.AccessLimit
	// 使用标准的 payload
	jwt.StandardClaims
}

type UserToken struct {}


// 将 JWT 字符串解密为 CustomClaims 对象
func (srv *UserToken) Decode(tokenStr string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return userConf.PrivateKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// 将 User 用户信息加密为 JWT 字符串
func (srv *UserToken) Encode(user *auth.LoginUser,limit *auth.AccessLimit) (string, error) {
	// 三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := UserClaims{
		user,
		limit,
		jwt.StandardClaims{
			Issuer:   userConf.Issuer, // 签发者
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(userConf.PrivateKey)
}

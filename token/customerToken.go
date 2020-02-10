package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/geiqin/supports/auth"
	"time"
)

type CustomerAble interface {
	Decode(tokenStr string) (*CustomerClaims, error)
	Encode(user *auth.LoginCustomer) (string, error)
}

// 自定义的 metadata，在加密后作为 JWT 的第二部分返回给客户端
type CustomerClaims struct {
	Customer *auth.LoginCustomer
	Limit *auth.AccessLimit
	// 使用标准的 payload
	jwt.StandardClaims
}

type CustomerToken struct {}


// 将 JWT 字符串解密为 CustomClaims 对象
func (srv *CustomerToken) Decode(tokenStr string) (*CustomerClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return customerConf.PrivateKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*CustomerClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// 将 User 用户信息加密为 JWT 字符串
func (srv *CustomerToken) Encode(customer *auth.LoginCustomer,limit *auth.AccessLimit) (string, error) {
	// 三天后过期
	expireTime := time.Now().Add(time.Hour * 24*3).Unix()
	claims := CustomerClaims{
		customer,
		limit,
		jwt.StandardClaims{
			Issuer:    customerConf.Issuer, // 签发者
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(customerConf.PrivateKey)
}

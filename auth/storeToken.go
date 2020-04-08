package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/geiqin/supports/helper/xtime"
)


type StoreAble interface {
	Decode(tokenStr string) (*StoreClaims, error)
	Encode(user *LoginStore) (string, error)
}

// 自定义的 metadata，在加密后作为 JWT 的第二部分返回给客户端
type StoreClaims struct {
	Store *LoginStore
	Limit *AccessLimit
	// 使用标准的 payload
	jwt.StandardClaims
}

type StoreToken struct {}


// 将 JWT 字符串解密为 CustomClaims 对象
func (srv *StoreToken) Decode(tokenStr string) (*StoreClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &StoreClaims{}, func(token *jwt.Token) (interface{}, error) {
		return storeConf.PrivateKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*StoreClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// 将 User 用户信息加密为 JWT 字符串
func (srv *StoreToken) Encode(store *LoginStore,limit *AccessLimit) (string, error) {
	expireTime :=xtime.GetAfterDay(storeConf.ExpireTime,xtime.DayType).Unix()
	claims := StoreClaims{
		store,
		limit,
		jwt.StandardClaims{
			Issuer:    storeConf.Issuer, // 签发者
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(storeConf.PrivateKey)
}

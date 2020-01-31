package token

import (
	"github.com/geiqin/supports/auth"
	"github.com/geiqin/supports/config"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// 定义加盐哈希密码时所用的盐，要保证其生成和保存都足够安全，比如使用 md5 来生成
var privateStoreKey []byte


type StoreAble interface {
	Decode(tokenStr string) (*StoreClaims, error)
	Encode(user *auth.CurrentStore) (string, error)
}

// 自定义的 metadata，在加密后作为 JWT 的第二部分返回给客户端
type StoreClaims struct {
	Store *auth.CurrentStore
	Limit *auth.AccessLimit
	// 使用标准的 payload
	jwt.StandardClaims
}

type StoreToken struct {}

func init() {
	keys :=config.GetConfig("store")
	prKey,ok :=keys["private_key"]
	if ok == false {
		log.Fatal("init private store key failed")
	}
	privateStoreKey =[]byte(prKey.(string))
}


// 将 JWT 字符串解密为 CustomClaims 对象
func (srv *StoreToken) Decode(tokenStr string) (*StoreClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &StoreClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateStoreKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*StoreClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// 将 User 用户信息加密为 JWT 字符串
func (srv *StoreToken) Encode(store *auth.CurrentStore,limit *auth.AccessLimit) (string, error) {
	// 三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := StoreClaims{
		store,
		limit,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.uim", // 签发者
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateStoreKey)
}

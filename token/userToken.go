package token

import (
	"github.com/geiqin/supports/auth"
	"github.com/geiqin/supports/config"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// 定义加盐哈希密码时所用的盐，要保证其生成和保存都足够安全，比如使用 md5 来生成
var privateUserKey []byte

type UserAble interface {
	Decode(tokenStr string) (*UserClaims, error)
	Encode(user *auth.CurrentUser) (string, error)
}


// 自定义的 metadata，在加密后作为 JWT 的第二部分返回给客户端
type UserClaims struct {
	User *auth.CurrentUser
	Limit *auth.AccessLimit
	// 使用标准的 payload
	jwt.StandardClaims
}

type UserToken struct {}

func init() {
	keys :=config.GetConfig("user")
	prKey,ok :=keys["private_key"]
	if ok == false {
		log.Fatal("init private user key failed")
	}
	privateUserKey =[]byte(prKey.(string))
}

// 将 JWT 字符串解密为 CustomClaims 对象
func (srv *UserToken) Decode(tokenStr string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateUserKey, nil
	})
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*UserClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// 将 User 用户信息加密为 JWT 字符串
func (srv *UserToken) Encode(user *auth.CurrentUser,limit *auth.AccessLimit) (string, error) {
	// 三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := UserClaims{
		user,
		limit,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.uim", // 签发者
			ExpiresAt: expireTime,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateUserKey)
}

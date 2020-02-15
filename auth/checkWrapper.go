package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"log"
)


// AuthWrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func CheckWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		userClaimVal:= meta["User-Claim"]
		storeClaimVal := meta["Store-Claim"]

		log.Println("userclaim:",userClaimVal)
		log.Println("storeclaim:",storeClaimVal)

		//用户授权
		if userClaimVal !=""{
			userClaim :=&UserClaims{}
			err:=json.Unmarshal([]byte(userClaimVal),userClaim)
			if err ==nil{
				log.Println("user authorization is ok")
				UserAuthorization(userClaim.User)
			}
		}

		//店铺授权
		if storeClaimVal !=""{
			storeClaim :=&StoreClaims{}
			err:=json.Unmarshal([]byte(storeClaimVal),storeClaim)
			if err==nil{
				log.Println("store authorization is ok")
				StoreAuthorization(storeClaim.Store)
			}
		}

		//继续执行下一步处理
		err := fn(ctx, req, resp)
		return err
	}
}

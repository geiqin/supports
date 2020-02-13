package auth

import (
	"context"
	"errors"
	//"github.com/geiqin/supports/cache"
	//"github.com/geiqin/supports/helper"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
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

		storeClaim := meta["User-Claim"]
		userClaim := meta["Store-Claim"]
		log.Println("userclaim:",userClaim)
		log.Println("storeclaim:",storeClaim)

		return nil
		//return
		/*
		mycache :=cache.GetCache()

		if storeAccessId != "" {
			log.Println("check wrapper store session id :",storeAccessId)
			store :=&LoginStore{}
			storeStr :=mycache.Get(storeAccessId)
			helper.JsonDecode(storeStr,store)
			StoreAuthorization(store)
		}

		if userAccessId != ""{
			log.Println("check wrapper user session id :",userAccessId)
			user :=&LoginUser{}
			userStr :=mycache.Get(userAccessId)
			helper.JsonDecode(userStr,user)
			UserAuthorization(user)
		}


		err := fn(ctx, req, resp)
		return err
		 */
	}
}

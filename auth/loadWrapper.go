package auth

import (
	"context"
	"errors"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
)

// AuthWrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func LoadWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}

		userId := meta["Auth-User-Id"]
		storeId := meta["Auth-Store-Id"]
		customerId := meta["Auth-Customer-Id"]

		if userId != "" {
			ctx = context.WithValue(ctx, "user_id", userId)
		}
		if storeId != "" {
			ctx = context.WithValue(ctx, "store_id", storeId)
		}
		if customerId != "" {
			ctx = context.WithValue(ctx, "customer_id", customerId)
		}

		//继续执行下一步处理
		err := fn(ctx, req, resp)
		return err
	}
}

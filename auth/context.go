package auth

import (
	"context"
	"github.com/geiqin/supports/helper"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/metadata"
	"net/http"
)

func StoreContext(storeId int64) context.Context {
	storeIdStr := helper.Int64ToString(storeId)
	ctx := context.WithValue(context.Background(), "store_id", storeIdStr)
	ctx = metadata.NewContext(ctx, map[string]string{
		"Auth-Store-Id": storeIdStr,
	})
	return ctx
}

func StoreContextByString(storeId string) context.Context {
	ctx := context.WithValue(context.Background(), "store_id", storeId)
	ctx = metadata.NewContext(ctx, map[string]string{
		"Auth-Store-Id": storeId,
	})
	return ctx
}

func StoreContextByBroker(p broker.Event) context.Context {
	if p != nil && p.Message().Header != nil {
		storeId := p.Message().Header["store_id"]
		if storeId != "" {
			sid := helper.StringToInt64(storeId)
			if sid > 0 {
				ctx := context.WithValue(context.Background(), "store_id", storeId)
				ctx = metadata.NewContext(ctx, map[string]string{
					"Auth-Store-Id": storeId,
				})
				return ctx
			}
		}
	}
	return context.Background()
}

//主要提供给 WEB模式下使用
func StoreContextByHeader(header http.Header) context.Context {
	storeId := header.Get("Auth-Store-Id")
	userId := header.Get("Auth-User-Id")
	customerId := header.Get("Auth-Customer-Id")
	ctx := context.Background()
	if storeId != "" {
		ctx = context.WithValue(ctx, "store_id", storeId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Store-Id": storeId,
		})
	}
	if userId != "" {
		ctx = context.WithValue(ctx, "user_id", userId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-User-Id": userId,
		})
	}
	if customerId != "" {
		ctx = context.WithValue(ctx, "customer_id", customerId)
		ctx = metadata.NewContext(ctx, map[string]string{
			"Auth-Customer-Id": userId,
		})
	}
	return ctx
}

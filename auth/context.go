package auth

import (
	"context"
	"github.com/geiqin/supports/helper"
	"github.com/micro/go-micro/broker"
)

func StoreContext(storeId int64) context.Context {
	ctx := context.WithValue(context.Background(), "store_id", helper.Int64ToString(storeId))
	return ctx
}

func StoreContextByString(storeId string) context.Context {
	ctx := context.WithValue(context.Background(), "store_id", storeId)
	return ctx
}

func StoreContextByBroker(broker broker.Event) context.Context {
	if broker != nil && broker.Message().Header != nil {
		storeId := broker.Message().Header["store_id"]
		if storeId != "" {
			sid := helper.StringToInt64(storeId)
			if sid > 0 {
				ctx := context.WithValue(context.Background(), "store_id", storeId)
				return ctx
			}
		}
	}
	return context.Background()
}

package auth

import (
	"context"
	"github.com/geiqin/supports/helper"
)

func StoreContext(storeId int64) context.Context {
	ctx := context.WithValue(context.Background(), "store_id", helper.Int64ToString(storeId))
	return ctx
}

func StoreContextByString(storeId string) context.Context {
	ctx := context.WithValue(context.Background(), "store_id", storeId)
	return ctx
}

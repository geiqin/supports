package auth

import (
	"context"
	"github.com/geiqin/supports/helper"
)

//获得当前用户ID
func GetUserId(ctx context.Context) int64 {
	val := ctx.Value("user_id")
	if val != nil {
		v := helper.StringToInt64(val.(string))
		return v
	}
	return 0
}

//获得当前店铺ID
func GetStoreId(ctx context.Context) int64 {
	val := ctx.Value("store_id")
	if val != nil {
		v := helper.StringToInt64(val.(string))
		return v
	}
	return 0
}

//获得当前客户ID
func GetCustomerId(ctx context.Context) int64 {
	val := ctx.Value("customer_id")
	if val != nil {
		v := helper.StringToInt64(val.(string))
		return v
	}
	return 0
}

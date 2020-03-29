package auth

import "context"

func StoreContext(storeId string) context.Context {
	ctx := context.WithValue(context.Background(), "store_id", storeId)
	return ctx
}

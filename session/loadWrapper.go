package session

import (
	"context"
	"errors"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
)


func LoadWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no session meta-data found in request")
		}

		log.Println("load session wrapper")
		sessId :=meta["Session-Id"]
		if sessId !=""{
			Start(sessId)
		}

		ctx =context.WithValue(ctx, "store_id", 5)

		err := fn(ctx, req, resp)
		return err
	}
}

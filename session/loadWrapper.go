package session

import (
	"context"
	"errors"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
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
		err := fn(ctx, req, resp)
		return err
	}
}

package session

import (
	//"geiqin.srv.uim/lib/auth"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
	"log"
	"context"
	"errors"
)


func LoadWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no session meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		sid := meta["SessionId"]

		SessionStart(sid)

		log.Println("session id:", sid)
		//ctx =context.WithValue(ctx,"userid",ausr)
		err := fn(ctx, req, resp)
		return err
		//return err
	}
}

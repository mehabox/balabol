package http

import (
	"time"

	"github.com/mehabox/balabol"
	routing "github.com/qiangxue/fasthttp-routing"
)

// logging is middleware for logging requests. Should be first in the chain for proper display.
func logging(next routing.Handler, logger balabol.Logger) routing.Handler {
	return func(ctx *routing.Context) error {
		t1 := time.Now()
		err := next(ctx)
		elapsed := time.Since(t1)
		logger.Printf(`[%15s] %s "%s" %d %s`, ctx.RemoteIP(), ctx.Method(), ctx.RequestURI(), ctx.Response.StatusCode(), elapsed)

		return err
	}
}

// slow is a useless middleware that just slows down response time more than a second.
func slow(next routing.Handler) routing.Handler {
	return func(ctx *routing.Context) error {
		time.Sleep(1 * time.Second)
		return next(ctx)
	}
}

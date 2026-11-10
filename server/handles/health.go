// Package handles contains HTTP request handlers.
package handles

import "github.com/valyala/fasthttp"

// Health responds with 200 OK. Used for liveness and readiness checks.
func Health(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString("ok")
}

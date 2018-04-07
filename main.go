package main

import (
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	cmd := parseCmd()
	fasthttp.ListenAndServe(cmd.addr, func(ctx *fasthttp.RequestCtx) {
		if cmd.debug {
			ctx.Request.Header.VisitAll(func(key, value []byte) {
				log.Printf("%s: %s", key, value)
			})
		}
		switch string(ctx.Path()) {
		case "/generate_204":
			ctx.SetStatusCode(204)
		default:
			ctx.Response.Header.SetContentType("text/plain; charset=utf-8")
			if cmd.ipHttpHeader == "NOT" {
				ctx.WriteString(ctx.RemoteIP().String())
			} else {
				ctx.Write(ctx.Request.Header.Peek(cmd.ipHttpHeader))
			}
		}
	})
}

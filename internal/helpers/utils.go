package helpers

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func JsonWrite(ctx *fasthttp.RequestCtx, code int, obj interface{}) error {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(code)
	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return err
	}
	return nil
}

func JsonRead(ctx *fasthttp.RequestCtx, obj interface{}) error {
	if err := json.Unmarshal(ctx.PostBody(), obj); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return err
	}
	return nil
}

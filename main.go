package main

import (
	"fmt"
	"github.com/Generalzy/General/General"
	"log"
	"net/http"
)

func LogMiddleWare()General.HandlerFunc{
	return func(ctx *General.Context) {
		fmt.Println("request",ctx.Method,ctx.Request.RemoteAddr,ctx.Path)
		ctx.Next()
		fmt.Println("response",ctx.Method,ctx.Request.RemoteAddr,ctx.Path)
	}
}

func main() {
	engine:=General.New()
	engine.Get("/", func(ctx *General.Context) {
		data:=General.H{}
		for k, v := range ctx.Request.Header{
			key:=fmt.Sprintf("Header[%q] = ",k)
			val:=fmt.Sprintf("%q \n",v)
			data[key]=val
		}
		ctx.Json(http.StatusOK,General.H{
			"code":0,
			"data":data,
			"err":"",
		})
	})
	engine.Get("/hello/:name", func(ctx *General.Context) {
		ctx.Json(http.StatusOK,General.H{
			"code":0,"data":ctx.Param("name"),"err":"",
		})
	})


	group:=engine.Group("/api/v1",LogMiddleWare())
	group.Get("/user", func(ctx *General.Context) {
		ctx.Json(http.StatusOK,General.H{
			"code":0,"data":ctx.Path,"err":"",
		})
	})
	log.Fatalln(engine.Run(":8080"))
}

package main

import (
	"fmt"
	"github.com/Generalzy/General/General"
	"log"
	"net/http"
)

func main() {
	engine:=General.Default()
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

	group:=engine.Group("/api/v1")
	group.Get("/panic", func(ctx *General.Context) {
		names := []string{"General_zy"}
		ctx.String(http.StatusOK, names[100])
	})

	log.Fatalln(engine.Run(":8080"))
}

package General

import (
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter()*router{
	return &router{handlers: make(map[string]HandlerFunc)}
}

// Url 用于注册视图函数与url的映射,灵感来自django
func (r *router)Url(method string,pattern string,handler HandlerFunc){
	k:=method+keyword+pattern
	r.handlers[k]=handler
}

// Path 等效于Url
func (r *router)Path(method string,pattern string,handler HandlerFunc){
	r.Url(method,pattern,handler)
}

// handle 路由映射
func (r *router)handle(ctx *Context){
	k:=ctx.Method+keyword+ctx.Path
	if handler,ok:=r.handlers[k];ok{
		handler(ctx)
	}else{
		ctx.String(http.StatusBadRequest,"404 not found: %s \n",ctx.Path)
	}
}


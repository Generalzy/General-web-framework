package General

import (
	"net/http"
	"strings"
)

const buf = 1<<2

type HttpMethod = string

type router struct {
	roots map[HttpMethod]*node
	handlers map[string]HandlerFunc
}

func NewRouter()*router{
	return &router{
		handlers: make(map[string]HandlerFunc,buf),
		roots: make(map[HttpMethod]*node,buf),
	}
}

// parsePattern 将Url按 / 切分
func parsePattern(pattern string)[]string{
	parts:=strings.Split(pattern,"/")
	newParts:=make([]string,0,buf)

	for _,part:=range parts{
		if part!=""{
			newParts=append(newParts,part)
		}
	}
	return newParts
}

// Url 用于注册视图函数与url的映射,灵感来自django
func (r *router)Url(method HttpMethod,pattern string,handler HandlerFunc){
	parts:=parsePattern(pattern)

	k:=method+keyword+pattern
	// 一个方法建立一个前缀树
	// 目前只有get 和 post 两个前缀树
	_,ok:=r.roots[method]
	if !ok{
		r.roots[method]= &node{}
	}
	// 向root插入路由
	r.roots[method].insert(pattern,parts,0)
	r.handlers[k]=handler
}

// getUrlParams 将路由中的动态参数导出
//
func (r *router)getUrlParams(method HttpMethod,path string)(*node,map[string]string){
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
		}
		return n, params
	}
	return nil, nil
}

// handle 路由映射
func (r *router)handle(ctx *Context){
	n,params:=r.getUrlParams(ctx.Method,ctx.Path)
	if n!=nil{
		ctx.Params=params
		k:=ctx.Method+keyword+n.pattern
		// node存在说明一定有该路由
		r.handlers[k](ctx)
	} else{
		ctx.String(http.StatusBadRequest,"404 not found: %v \n",ctx.Path)
	}
}


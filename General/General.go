package General

import (
	"fmt"
	"net/http"
)

const (
	MethodGet     = http.MethodGet
	MethodHead    = http.MethodHead
	MethodPost    = http.MethodPost
	MethodPut     = http.MethodPut
	MethodPatch   = http.MethodPatch
	MethodDelete  = http.MethodDelete
	MethodConnect = http.MethodConnect
	MethodOptions = http.MethodOptions
	MethodTrace   = http.MethodTrace
)

const keyword = "_"

// HandlerFunc 定义视图函数
type HandlerFunc func(w http.ResponseWriter,request *http.Request)

// Engine 定义General引擎
type Engine struct {
	router map[string]HandlerFunc
}

// New 初始化引擎
func New()*Engine{
	return &Engine{router: make(map[string]HandlerFunc)}
}

// Url 用于注册视图函数与url的映射,灵感来自django
func (e *Engine)Url(method string,pattern string,handler HandlerFunc){
	k:=method+keyword+pattern
	e.router[k]=handler
}

// Path 等效于Url
func (e *Engine)Path(method string,pattern string,handler HandlerFunc){
	e.Url(method,pattern,handler)
}

// Get HTTP GET请求
func (e *Engine)Get(pattern string,handler HandlerFunc){
	e.Url(MethodGet,pattern,handler)
}

// Post HTTP POST请求
func (e *Engine)Post(pattern string,handler HandlerFunc){
	e.Url(MethodPost,pattern,handler)
}

// ServeHTTP 实现Handler接口
func (e *Engine)ServeHTTP(w http.ResponseWriter,request *http.Request){
	k:=request.Method+keyword+request.URL.Path
	if handler,ok:=e.router[k];ok{
		handler(w,request)
	}else{
		_,_=fmt.Fprintf(w,"404 not found: %s \n",request.URL)
	}
}

// Run 开启HTTP服务
func (e *Engine)Run(addr string)error{
	return http.ListenAndServe(addr,e)
}
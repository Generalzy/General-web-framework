package General

import (
	"net/http"
)

const keyword = "_"

// HandlerFunc 定义视图函数
type HandlerFunc func(*Context)

// Engine 定义General引擎
type Engine struct {
	router *router
}

// New 初始化引擎
func New()*Engine{
	return &Engine{router: NewRouter()}
}

// Url 用于注册视图函数与url的映射,灵感来自django
func (e *Engine)Url(method string,pattern string,handler HandlerFunc){
	e.router.Url(method,pattern,handler)
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
	e.router.handle(newContext(w,request))
}

// Run 开启HTTP服务
func (e *Engine)Run(addr string)error{
	return http.ListenAndServe(addr,e)
}
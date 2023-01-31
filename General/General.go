package General

import (
	"net/http"
	"strings"
)

const keyword = "_"

// HandlerFunc 定义视图函数
type HandlerFunc func(*Context)

// RouterGroup 路由分组
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (g *RouterGroup)Group(prefix string,middleware...HandlerFunc)*RouterGroup{
	engine:=g.engine
	engine.groups = append(engine.groups,g)

	rg:= &RouterGroup{
		engine: engine,
		// 支持分组的分组...
		prefix: g.prefix+prefix,
		// 谁调用Group,就将谁设置为parent
		parent: g,
		// 将middleware加入
		middlewares: middleware,
	}
	// 把新的group也加进去
	engine.groups = append(engine.groups,rg)
	return rg
}

// Use 将中间件添加入路由
func (g *RouterGroup)Use(middlewares... HandlerFunc){
	g.middlewares=append(g.middlewares,middlewares...)
}

// Engine 定义General引擎
// 将引擎看作最大的Group
type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New 初始化引擎
func New()*Engine{
	engine:=&Engine{router: NewRouter()}
	// 引擎的父节点,代表引擎是root
	// 引擎的engine自然为自己
	// 引擎的前缀自然为""
	engine.RouterGroup = &RouterGroup{engine: engine}
	// 将引擎加入groups
	engine.groups=[]*RouterGroup{engine.RouterGroup}
	return engine
}

// Url 用于注册视图函数与url的映射,灵感来自django
func (g *RouterGroup)Url(method string,pattern string,handler HandlerFunc){
	pattern = g.prefix+pattern
	g.engine.router.Url(method,pattern,handler)
}

// Path 等效于Url
func (g *RouterGroup)Path(method string,pattern string,handler HandlerFunc){
	g.Url(method,pattern,handler)
}

// Get HTTP GET请求
func (g *RouterGroup)Get(pattern string,handler HandlerFunc){
	g.Url(MethodGet,pattern,handler)
}

// Post HTTP POST请求
func (g *RouterGroup)Post(pattern string,handler HandlerFunc){
	g.Url(MethodPost,pattern,handler)
}

// ServeHTTP 实现Handler接口
func (e *Engine)ServeHTTP(w http.ResponseWriter,request *http.Request){
	middlewares:=make([]HandlerFunc,0,buf)
	// 遍历路由组
	for _,group:=range e.groups{
		// 如果当前路由是group定义的前缀
		if strings.HasPrefix(request.URL.Path,group.prefix){
			// 将应用到group的middlewares取出来
			middlewares=append(middlewares,group.middlewares...)
		}
	}
	ctx:=newContext(w,request)
	// 将Group的middleware交给ctx
	ctx.middlewares=middlewares
	// handle 路由映射
	e.router.handle(ctx)
}

// Run 开启HTTP服务
func (e *Engine)Run(addr string)error{
	return http.ListenAndServe(addr,e)
}
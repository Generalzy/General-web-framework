package General

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type any interface {}

type H map[string]any

type Context struct {
	Request *http.Request
	Writer http.ResponseWriter

	Path string
	Method string
	Status int
}

func newContext(w http.ResponseWriter, request *http.Request)*Context{
	return &Context{
		Path: request.URL.Path,
		Method: request.Method,
		Request: request,
		Writer: w,
	}
}

// SetHeader 设置响应头信息
func (c *Context)SetHeader(key string,val string){
	c.Writer.Header().Set(key,val)
}

// SetStatus 设置响应状态码
func (c *Context)SetStatus(code int){
	c.Status=code
	c.Writer.WriteHeader(code)
}

// String 响应字符格式的快捷操作
func (c *Context)String(code int,format string,value...any){
	c.SetHeader("Content-Type",ContentText)
	c.SetStatus(code)
	_,_ = c.Writer.Write([]byte(fmt.Sprintf(format,value)))
}

// Json 响应json格式的快捷操作
func (c *Context)Json(code int,obj H){
	c.SetHeader("Content-Type",ContentJson)
	c.SetStatus(code)
	data,err:=json.Marshal(obj)
	if err!=nil{
		http.Error(c.Writer,err.Error(),http.StatusInternalServerError)
	}else{
		_,_ = c.Writer.Write(data)
	}
}

// HTML 响应html格式的快捷操作
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	_,_ = c.Writer.Write([]byte(html))
}




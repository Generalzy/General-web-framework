package General

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// trace 获取触发 panic 的堆栈信息
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.String(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}

func Logger()HandlerFunc{
	// 2023/01/31 15:20:23  [GET] [500] /panic in 340.1µs
	return func(ctx *Context) {
		start:=time.Now()
		ctx.Next()
		log.Printf(" [%s] [%d] %s in %v \n",ctx.Method,ctx.Status,ctx.Path,time.Since(start))
	}
}

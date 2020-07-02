package fasthttpserv


import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/rz1226/blackboardkit"
	"github.com/valyala/fasthttp"
	"log"
)

const RETURN_PANIC = `{"code":-2,"msg":"fast-h server error"}`

type routerFunc func(*fasthttp.RequestCtx) string

type FastHTTPServ struct {
	router *fasthttprouter.Router
}

func NewServ() FastHTTPServ {
	res := FastHTTPServ{}
	res.router = fasthttprouter.New()
	return res
}

func (fhs FastHTTPServ) GET(path string, f routerFunc) {
	fhs.router.GET(path, makeApifunc(f, path))
}

func (fhs FastHTTPServ) POST(path string, f routerFunc) {
	fhs.router.POST(path, makeApifunc(f, path))
}

func (fhs FastHTTPServ) HEAD(path string, f routerFunc) {
	fhs.router.POST(path, makeApifunc(f, path))
}

func (fhs FastHTTPServ) OPTIONS(path string, f routerFunc) {
	fhs.router.POST(path, makeApifunc(f, path))
}

func (fhs FastHTTPServ) PUT(path string, f routerFunc) {
	fhs.router.POST(path, makeApifunc(f, path))
}

func (fhs FastHTTPServ) PATCH(path string, f routerFunc) {
	fhs.router.POST(path, makeApifunc(f, path))
}

func (fhs FastHTTPServ) DELETE(path string, f routerFunc) {
	fhs.router.POST(path, makeApifunc(f, path))
}

func (fhs FastHTTPServ) Start(port string) {

	log.Fatal(fasthttp.ListenAndServe(":"+port, fhs.router.Handler))

}

// 第二个参数path用于日志
func makeApifunc(f routerFunc, path string) fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if co := recover(); co != nil {
				blackboardkit.DefaultBB.API.Panic("panic with api path:", path, co)
				fmt.Fprint(ctx, RETURN_PANIC, co, "api path =", path)
			}
		}()
		t := blackboardkit.DefaultBB.API.Start("访问" + path + "接口")
		str := f(ctx)
		blackboardkit.DefaultBB.API.End(t)
		fmt.Fprint(ctx, str)
	}
}





package base

import (
	"github.com/astaxie/beego/context"
	"net/http"
	"path"
	"strings"
)

const staticFileExt = ".html|.js|.css|.png|.jpg|.jpeg|.ico|.otf"

func WebServerFilter(ctx *context.Context) {
	urlPath := ctx.Request.URL.Path
	if urlPath == "" || urlPath == "/" {
		urlPath = "index.html"
	}

	ext := path.Ext(urlPath) // 后缀名

	if ext == "" {
		return
	}

	index := strings.Index(staticFileExt, ext)
	if index >= 0 {
		http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+urlPath)
	}
}

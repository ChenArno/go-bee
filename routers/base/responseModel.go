package base

import (
	"compress/gzip"
	"encoding/json"
	"github.com/astaxie/beego"
	"io"
	"strings"
)

type ResponseModel struct {
	beego.Controller
	response response
}

// 定义结构体
type response struct {
	Code         int         `json:"code"`
	Msg          string      `json:"msg"`
	ResultStatus bool        `json:"resultStatus"`
	Data         interface{} `json:"data"`
}

// 定义方法
//response格式化
func (r *ResponseModel) ResponseResult(errorNo Errno, data interface{}) {
	jsonData := response{
		Code:         errorNo.Code,
		Msg:          errorNo.Msg,
		ResultStatus: errorNo.resultStatus,
		Data:         data,
	}

	returnJSON, err := json.Marshal(jsonData)

	if err != nil {
		beego.Error(err)
	}
	r.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	//使用gzip原始，json数据会只有原本数据的10分之一左右
	if strings.Contains(strings.ToLower(r.Ctx.Request.Header.Get("Accept-Encoding")), "gzip") {
		r.Ctx.ResponseWriter.Header().Set("Content-Encoding", "gzip")
		//gzip压缩
		w := gzip.NewWriter(r.Ctx.ResponseWriter)
		defer w.Close()
		w.Write(returnJSON)
		w.Flush()
	} else {
		io.WriteString(r.Ctx.ResponseWriter, string(returnJSON))
	}
	r.StopRun()
}

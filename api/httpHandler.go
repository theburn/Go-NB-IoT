package api

import (
	"fmt"
	"html/template"
	"io/ioutil"

	"github.com/theburn/Go-NB-IoT/configure"
	log "github.com/theburn/Go-NB-IoT/logging"

	"github.com/valyala/fasthttp"
)

type deviceServiceData struct {
	ServiceID   string      `json:"serviceId"`
	ServiceType string      `json:"serviceType"`
	Data        interface{} `json:"data"`
	EventTime   string      `json:"eventTime"`
}

type deviceDataChanged struct {
	NotifyType string            `json:"notifyType"`
	RequestID  string            `json:"requestId"`
	DeviceID   string            `json:"deviceId"`
	GatewayId  string            `json:"gatewayId"`
	Service    deviceServiceData `json:"service"`
}

type IHandle interface {
	NotifyHandler(notifyType string, postBody []byte) error
}

var DoHandle IHandle

func InitDoHandle(i IHandle) {
	DoHandle = i
}

func CallBackHandler(ctx *fasthttp.RequestCtx) {
	log.Debug(">>>> String", string(ctx.PostBody()))
	v, ok := ctx.UserValue("subcribeNotifyType").(string)
	if !ok {
		ctx.SetStatusCode(500)
		fmt.Fprint(ctx)
		return
	}

	if err := DoHandle.NotifyHandler(v, ctx.PostBody()); err != nil {
		log.Errorf("NotifyHandler error:", err.Error())
		ctx.SetStatusCode(500)
	} else {
		ctx.SetStatusCode(200)
	}

	fmt.Fprint(ctx)
	return
}

func GetServerLogs(ctx *fasthttp.RequestCtx) {
	body, _ := ioutil.ReadFile("logs/notice.log") // temp handler TODO

	fmt.Fprint(ctx, string(body))

}

//----------------- status ------------------

func ServStatic(ctx *fasthttp.RequestCtx) {
	log.Debug(">> static: %s", ctx.Path())
	fsHandler(ctx)
}

func StatusHandlerV1(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")
	t, _ := template.ParseFiles(configure.NBIoTConfig.ServerParam.StaticPath + "/index.html")
	t.Execute(ctx, "")
}

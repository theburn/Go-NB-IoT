package api

import (
	"Go-NB-IoT/amqp"
	"Go-NB-IoT/configure"
	log "Go-NB-IoT/logging"
	"fmt"
	"html/template"
	"io/ioutil"

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

func CallBackDeviceDataChanged(ctx *fasthttp.RequestCtx) {

	//var httpDeviceDataChanged deviceDataChanged

	log.Debug(">>>> String", string(ctx.PostBody()))

	/*
		// Telecom Test
			if string(ctx.PostBody()) == "push success." {
				log.Info("Test Push Success! ")
				ctx.SetStatusCode(200)
			} else {
				amqp.AMQPSend(amqp.DefaultQueueName, amqp.DefaultContentType, ctx.PostBody())
				ctx.SetStatusCode(500)
			}
	*/

	if err := amqp.AMQPSend(amqp.DefaultQueueName, amqp.DefaultContentType, ctx.PostBody()); err != nil {
		ctx.SetStatusCode(200)
	} else {
		ctx.SetStatusCode(500)
	}

	fmt.Fprint(ctx)
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

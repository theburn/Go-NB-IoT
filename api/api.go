package api

import (
	log "github.com/theburn/Go-NB-IoT/logging"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var apiRouter *fasthttprouter.Router
var fsHandler func(ctx *fasthttp.RequestCtx)

func init() {
	apiRouter = fasthttprouter.New()
}

func Run(listenPort, staticPath string) error {
	log.Info("Start HTTP ListenAndServe ... ")
	// request /static/css/xxx.css -> css/xxx.css
	fsHandler = fasthttp.FSHandler(staticPath, 1)
	ListenPort := ":" + listenPort

	apiRouter.POST("/api/callback/v1.5.1/deviceDataChanged", CallBackDeviceDataChanged)
	apiRouter.GET("/static/*filepath", ServStatic)
	apiRouter.GET("/logs", GetServerLogs)

	if err := fasthttp.ListenAndServe(ListenPort, apiRouter.Handler); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

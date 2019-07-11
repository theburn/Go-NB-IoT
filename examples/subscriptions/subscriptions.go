package main

import (
	"encoding/json"
	"flag"
	"time"

	"github.com/theburn/Go-NB-IoT/client"
	"github.com/theburn/Go-NB-IoT/configure"
	log "github.com/theburn/Go-NB-IoT/logging"
	"github.com/theburn/Go-NB-IoT/subscriptions"
	"github.com/theburn/Go-NB-IoT/utils"
)

var (
	version        = "v1.0"
	configfilename = flag.String("cfgfile", "conf/config.json", "live configure filename")
	loglevel       = flag.String("loglevel", "info", "log level")
	logfile        = flag.String("logfile", "logs/go-nb-iot.log", "log file path")
)

func init() {
	flag.Parse()
	log.SetOutputByName(*logfile)
	log.SetRotateByDay()
	log.SetLevelByString(*loglevel)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Error("go-nb-iot panic: ", r)
			time.Sleep(1 * time.Second)
		}
	}()
	err := configure.LoadConfig(*configfilename)
	if err != nil {
		return
	}

	log.SetMaxLogDay(7)

	// output system info
	log.Info("-----------------START----------------")
	log.Info("start go-nb-iot: ", version)

	s := subscriptions.SubscriptionsBusinessReq{}
	s.AppId = configure.NBIoTConfig.ReqParam.AppID
	s.CallbackUrl = "https://test-api.lng-tech.net:9999/api/callback/v1.5.1/deviceDataChanged"
	s.NotifyType = subscriptions.NodeTypeDict.DeviceDatasChanged

	if c, err := client.NewNBHttpClient(); err != nil {
		log.Error("New Client error", err)
	} else {
		c.Login()
		resp, _ := s.SubscriptionsBusiness(c)
		data, _ := json.Marshal(resp)
		utils.LogNoticeToFile(string(data))

		log.Debugf("%+v", *resp)

		resp2, _ := resp.SubscriptionsQuerySingle(c)
		data, _ = json.Marshal(resp2)
		utils.LogNoticeToFile(string(data))

		log.Debugf("%+v", *resp2)

	}

}

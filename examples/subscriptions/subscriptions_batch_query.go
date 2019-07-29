package main

import (
	"encoding/json"
	"flag"
	"fmt"
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
	notifyType     = flag.String("notifyType", "", "notifyType")
)

/*
1.bindDevice（绑定设备，订阅后推送绑定设备通知）
2.deviceAdded（添加新设备，订阅后推送注册设备通知）
3.deviceInfoChanged（设备信息变化，订阅后推送设备信息变化通知）
4.deviceDataChanged（设备数据变化，订阅后推送设备数据变化通知）
5.deviceDatasChanged（设备数据批量变化，订阅后推送批量设备数据变化通知）
6.deviceDeleted（删除设备，订阅后推送删除设备通知）
7.serviceInfoChanged（服务信息变化，订阅后推送设备服务信息变化通知）
8.ruleEvent（规则事件，订阅后推送规则事件通知）
9.deviceModelAdded（添加设备模型，订阅后推送增加设备模型通知）
10.deviceModelDeleted（删除设备模型，订阅后推送删除设备模型通知）
*/

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

	if c, err := client.NewNBHttpClient(); err != nil {
		log.Error("New Client error", err)
	} else {
		c.Login()
		resp2, _ := subscriptions.SubscriptionsQueryBatch(c, *notifyType)
		data, _ := json.Marshal(resp2)
		utils.LogNoticeToFile(string(data))

		fmt.Printf("data  : %+v\n", string(data))

	}

}

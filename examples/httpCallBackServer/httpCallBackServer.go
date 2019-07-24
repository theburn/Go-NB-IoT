package main

import (
	"flag"
	"time"

	"github.com/theburn/Go-NB-IoT/api"
	"github.com/theburn/Go-NB-IoT/configure"
	log "github.com/theburn/Go-NB-IoT/logging"
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

type notify struct{}

func (n notify) NotifyHandler(notifyType string, postBody []byte) error {
	return AMQPSend("BIoTCallback-Telecom", notifyType, postBody)
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

	if err := InitAMQP(configure.NBIoTConfig.AMQPParam.AMQPURL); err != nil {
		log.Errorf("init amqpQueue error: ", err.Error())
	}

	if err := InitQueue("BIoTCallback-Telecom"); err != nil {
		log.Errorf("init Queue error: ", err.Error())
	}

	N := notify{}

	api.InitDoHandle(N)

	log.Info("amqpQueue init success..!")

	api.Run("19999", "./static")

}

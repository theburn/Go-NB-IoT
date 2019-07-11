package main

import (
	"flag"
	"time"

	"github.com/theburn/Go-NB-IoT/amqpQueue"
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

	if err := amqpQueue.InitAMQP(configure.NBIoTConfig.AMQPParam.AMQPURL); err != nil {
		log.Errorf("amqpQueue init amqpQueue error: ", err.Error())
	}

	if err := amqpQueue.InitQueue(amqpQueue.DefaultQueueName); err != nil {
		log.Errorf("amqpQueue init Queue error: ", err.Error())
	}

	log.Info("amqpQueue init success..!")

	api.Run("9999", "./static")

}

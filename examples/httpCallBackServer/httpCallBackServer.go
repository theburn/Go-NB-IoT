package main

import (
	"flag"
	"time"

	"github.com/theburn/Go-NB-IoT/amqp"
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

	if err := amqp.InitAMQP(); err != nil {
		log.Errorf("amqp init amqp error: ", err.Error())
	}

	if err := amqp.InitQueue(amqp.DefaultQueueName); err != nil {
		log.Errorf("amqp init Queue error: ", err.Error())
	}

	log.Info("amqp init success..!")

	api.Run()

}
